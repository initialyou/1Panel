package components

type Block struct {
	Comment    string
	Directives []IDirective
}

func (b *Block) GetDirectives() []IDirective {
	return b.Directives
}

func (b *Block) GetComment() string {
	return b.Comment
}

func (b *Block) FindDirectives(directiveName string) []IDirective {
	directives := make([]IDirective, 0)
	for _, directive := range b.GetDirectives() {
		if directive.GetName() == directiveName {
			directives = append(directives, directive)
		}
		if directive.GetBlock() != nil {
			directives = append(directives, directive.GetBlock().FindDirectives(directiveName)...)
		}
	}

	return directives
}

func (b *Block) UpdateDirectives(directiveName string, directive Directive) {
	directives := b.GetDirectives()
	index := -1
	for i, dir := range directives {
		if dir.GetName() == directiveName {
			index = i
			break
		}
	}
	if index > -1 {
		directives[index] = &directive
	} else {
		directives = append(directives, &directive)
	}
	b.Directives = directives
}

func (b *Block) UpdateDirectiveBySecondKey(name string, key string, directive Directive) {

	directives := b.GetDirectives()

	index := -1
	for i, dir := range directives {
		if dir.GetName() == name && dir.GetParameters()[0] == key {
			index = i
			break
		}
	}
	if index > -1 {
		directives[index] = &directive
	} else {
		directives = append(directives, &directive)
	}
	b.Directives = directives
}

func (b *Block) AddDirectives(directive Directive) {
	directives := append(b.GetDirectives(), &directive)
	b.Directives = directives
}

func (b *Block) RemoveDirectives(names []string) {
	nameMaps := make(map[string]struct{}, len(names))
	for _, name := range names {
		nameMaps[name] = struct{}{}
	}
	directives := b.GetDirectives()
	var newDirectives []IDirective
	for _, dir := range directives {
		if _, ok := nameMaps[dir.GetName()]; ok {
			continue
		}
		newDirectives = append(newDirectives, dir)
	}
	b.Directives = newDirectives
}