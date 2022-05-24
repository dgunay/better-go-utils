package conversion

type From[Self any, Source any] interface {
	From(val Source) Self
}

type Into[Target any] interface {
	Into() Target
}
