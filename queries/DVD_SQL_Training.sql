select c.first_name, round(p.amount, 0) as estimated from payment p
join public.customer c on p.customer_id = c.customer_id
order by c.first_name desc;

select avg(p.amount) from public.payment p;

/*
adicionando com junção de tabelas que é mais eficiente
junta os dados FROM
realiza o filtro WHERE
agrupa os dados com o group
faz o select das colunas
ordena os dados com o order by

*/
select
    c.customer_id, concat(c.first_name,' ', c.last_name) as customer_name, count(p.payment_id) as payment_amount
from public.customer c
left join public.payment p on c.customer_id = p.customer_id
group by c.customer_id
order by 3 desc;

/*
agrupando com subselect.
para cada linha de dados a subconsulta é re-executada
*/
select
    c.customer_id,
    concat(c.first_name, ' ', c.last_name),
    (
        select count(p.payment_id) from public.payment p
        where p.customer_id = c.customer_id
                                   ) as payment_amount
    from public.customer c
group by c.customer_id
order by 3 desc
limit 10;

/*
filtra os usuários que só possuem 1 pagamento

count(distinct concat(c.first_name, ' ', c.last_name)) as unique
distinct c.customer_id, concat(c.first_name, ' ', c.last_name)
*/
select
    count(distinct concat(c.first_name, ' ', c.last_name)) as unique
from public.customer c;

/*
Buscar clientes que não possuem pagamento
*/

select c.customer_id, concat(c.first_name, ' ', c.last_name) as customer_name, count(p.payment_id) as payment_amount from customer c
left join public.payment p on c.customer_id = p.customer_id
group by c.customer_id, concat(c.first_name, ' ', c.last_name)
having count(p.payment_id) = 0;
-- nunca insira uma condição na claúsula having que não envolta agragação, essas codições
-- são mais eficiêntes na clausula WHERE.

/*
Subconsulta com with e CTE
A vantagem de se usar subconsulta é que podemos transformar problemas maiores em menores
neste exemplo estamos utilizando uma subconsulta em uma tabela derivada.
*/
select po.payment_id, c.first_name
from (select * from payment p where p.staff_id = 1) po
join customer c on po.customer_id = c.customer_id;

-- WITH cte (common table expression)
WITH po as (
    select * from payment p where p.staff_id = 1
)
select po.customer_id, c.first_name from  po
inner join customer c on po.customer_id = c.customer_id;

/*
GROUP BY
Permite que filtremos um determinado grupo de dados, para funcionar, no select precisamor definir qual a coluna de agrupamento
tipo um ID, ou NOME, casos em que podem aparecer mais de uma vez, depois disso uma função de agregação,
tipo o count(), e por fim caso queiramos filtrar esse dados,
podemos usar o Having
*/

select table_name from information_schema.views
where table_schema NOT IN('information_schema', 'pg_catalog');

select * from actor_info;

select 4-3;

-- Operações e funções.
select * from actor_info
where actor_id not between 1 and 10;

select * from actor_info
where film_info like '%Dance';

-- Use o exists quando quisermos que somente os valores de uma única tabela seja retornado.
-- As vezes esse tipo de consulta é chamado de semijunção.
-- Win
select c.first_name, c.last_name, c.active from customer c
where not exists(
    select 1 from payment p
    where c.customer_id = p.customer_id
);

select c.customer_id, c.first_name, c.last_name, c.active, count(p.payment_id) as total from customer c
left join payment p on p.customer_id = c.customer_id
group by c.customer_id
having count(p.payment_id) = 0;

select c.customer_id, c.first_name, c.last_name, c.active from customer c
left join payment p on p.customer_id = c.customer_id
where p.payment_id is null;

-- NOT IN for usado, se houver ao menos um valor null na subconsulta, esta nuca será igual true,
-- o que significa que o nenhuma linha será retornada.
-- opte pelo exists nesses casos.

-- Operador like
/*
% para um ou mais caracters
_ para um caracter em específico
*/
select * from actor_info
where film_info like '%Dance%';

select * from actor_info ai
where ai.first_name like '_ean';


-- Funções de agregação
select min(c.last_update), max(c.last_update) from customer c
group by c.first_name
having max(c.last_update) > '2013-10-22';



select p.customer_id,  p.amount, count(p.customer_id), extract(minute from p.payment_date) from payment p
where extract(minute from p.payment_date) < 10
group by p.customer_id, p.amount, p.payment_date
having count(p.customer_id) > 1;


select trunc(98.7654, 2);

-- diferença entre duas datas
select extract(minute from '2023-10-11 00:10:02'::timestamp - '2023-10-11 00:00:01'::timestamp);

-- busco os pagamentos com diferença de 2 a 4 minutos.
-- seed insert into payment (customer_id, staff_id, rental_id, amount, payment_date) values (263,1,15293,0.99,'2007-05-14 13:46:29.996577')
/*
Busco todos os pagamentos maior que data
Depois realiza a junção self considerando apenas os p2 que possuem o mesmo cliente id
filtra os pagamentos p2.id deferente de s.id linha a linha,
filtra os pagamentos que tem uma diferença de 1 a 4 minutos
agrupa para remover os duplicados
*/
select p2.customer_id, p2.payment_id, s.payment_id from (
    select p.payment_id, p.customer_id, p.payment_date from payment p
             where p.payment_date >= '2007-05-14' -- filtro os pagamentos maiores que a data tal
              ) s
join payment p2 on s.customer_id = p2.customer_id -- self join com mesmo cliente id
where p2.payment_id <> s.payment_id and extract(minute from p2.payment_date - s.payment_date) >= 1
and p2.payment_date >= '2007-05-14'
and extract(minute from p2.payment_date - s.payment_date) <= 4 -- busco pagamentos com diferença de 4 minutos
group by p2.customer_id, p2.payment_id, s.payment_id
order by p2.customer_id;

-- menor performance por nested loop
select p2.customer_id, p2.payment_id from payment p2
where p2.payment_date >= '2007-05-14' and exists(
    select p.payment_id, p.customer_id, p.payment_date from payment p
    where p.payment_date >= '2007-05-14'
    and p2.payment_id <> p.payment_id and extract(minute from p2.payment_date - p.payment_date) >= 1
    and extract(minute from p2.payment_date - p.payment_date) <= 4 -- busco pagamentos com diferença de 4 minutos
)
group by p2.customer_id, p2.payment_id;


select to_char(date '2023-03-16', 'day');

select date_trunc('month', date '2020-02-25');