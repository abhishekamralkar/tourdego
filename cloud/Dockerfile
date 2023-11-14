from golang:1.21 as base

workdir /work

from golang:1.21 as build

workdir /app
copy ./* /app/
run go build -o infra

from alpine as runner
copy --from=build /app/infra /
CMD ./infra 
