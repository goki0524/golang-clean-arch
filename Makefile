build:
	wire && go1.18 build server.go wire_gen.go
run:
	wire && go1.18 run server.go wire_gen.go
