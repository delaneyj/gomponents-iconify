package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackMg0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackMg1)"><use href="#flagpackMg0"/><path fill="#78D843" fill-rule="evenodd" d="M12 12h20v12H12V12Z" clip-rule="evenodd"/><path fill="#EA1A1A" fill-rule="evenodd" d="M12 0h20v12H12V0Z" clip-rule="evenodd"/><path fill="#F7FCFF" fill-rule="evenodd" d="M0 0h12v24H0V0Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackMg1"><use href="#flagpackMg0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}