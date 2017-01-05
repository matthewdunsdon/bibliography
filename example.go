package bibliography

type entity struct {
	aValue string `bibtexer:"ignore"`
	Entry
}

type AA func() (e entity)

func createEntity() (e entity) {
	e = entity{"abc", Entry{}}
	return
}
