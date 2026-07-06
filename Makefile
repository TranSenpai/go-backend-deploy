.PHONY: ent ent-new

ent:
	go generate ./ent

ent-new:
	go run -mod=mod entgo.io/ent/cmd/ent new ${name}


