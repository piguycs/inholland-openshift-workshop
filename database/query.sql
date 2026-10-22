-- name: GetAuthor :one
select * from authors
where id = $1 limit 1;

-- name: ListAuthors :many
select * from authors
order by name;

-- name: CreateAuthor :one
insert into authors (
    name, bio
) values (
    $1, $2
)
returning *;

-- name: UpdateAuthor :exec
update authors
set name = $2,
    bio = $3
where id = $1
returning *;

-- name: DeleteAuthor :exec
delete from authors
where id = $1;
