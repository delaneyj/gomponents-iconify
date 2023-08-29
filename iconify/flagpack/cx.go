package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackCx0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackCx1)"><use href="#flagpackCx0"/><path fill="#5EAA22" fill-rule="evenodd" d="M0 0h32v24H0V0Z" clip-rule="evenodd"/><path fill="#4141DB" fill-rule="evenodd" d="m0 0l32 24H0V0Z" clip-rule="evenodd"/><path fill="#fff" fill-rule="evenodd" d="m6 9l-1 .732l.134-1.232L4 8l1.134-.5L5 6.268L6 7l1-.732L6.866 7.5L8 8l-1.134.5L7 9.732L6 9Zm-2 6l-1 .732l.134-1.232L2 14l1.134-.5L3 12.268L4 13l1-.732l-.134 1.232L6 14l-1.134.5L5 15.732L4 15Zm4 6l-1 .732l.134-1.232L6 20l1.134-.5L7 18.268L8 19l1-.732l-.134 1.232L10 20l-1.134.5L9 21.732L8 21Zm1-9.5l-.5.366l.067-.616L8 11l.567-.25l-.067-.616l.5.366l.5-.366l-.067.616L10 11l-.567.25l.067.616L9 11.5Z" clip-rule="evenodd"/><path fill="#FECA00" fill-rule="evenodd" d="M21.561 12.224c-.374-.39 2.99-1.306 3.339-2.488c.462-1.029.153-1.862-1.593-2.594c-1.746-.732-3.51-2.148-1.481-2.148s2.029.29 2.574 1.143c.545.854 1.724 1.036 1.748 0c0-1.766.167-2.003-1.258-3.538c-1.426-1.536 3.574.63 3.16 3.262c-.412 2.632-.91 1.782-.594 2.224c.316.442 2.216-.93 1.979.897c-.665 1.051-1.984 2.614-7.874 3.242Z" clip-rule="evenodd"/><path fill="#FECA00" fill-rule="evenodd" d="M16 17a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z" clip-rule="evenodd"/><path fill="#548650" fill-rule="evenodd" d="M12.953 10.162c.567-.126 1.343 1.392 2.742.823c1.4-.57 1.692-1.657 2.357-1.26c.666.399.753 1.357.31 1.838c-.443.482-1.195.596-1.195 1.205c0 .609-.148 2.78-.435 2.15c-.861-.699-.336-1.812-1.397-2.15c-1.061-.338-1.706-.327-2.489-.184c-.783.143.526-.392.804-.9c.565-.543-.956-1.341-.697-1.522Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackCx1"><use href="#flagpackCx0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}