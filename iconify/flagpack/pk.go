package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackPk0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackPk1)"><use href="#flagpackPk0"/><path fill="#2F8D00" fill-rule="evenodd" d="M8 0h24v24H8V0Z" clip-rule="evenodd"/><path fill="#F7FCFF" fill-rule="evenodd" d="M0 0h8v24H0V0Z" clip-rule="evenodd"/><path fill="#F1F9FF" fill-rule="evenodd" d="M22.43 15.306s-4.466 1.165-8.011-1.21c-3.545-2.375-1.763-7.848-1.763-7.848c-1.849.269-4.752 7.015-.072 10.398c4.68 3.384 9.164.131 9.845-1.34Zm-4.956-6.44l-2.372 1.16l2.502.446l.338 2.445l1.417-2.083l2.788.189l-2.184-1.63l1.163-2.176l-2.174.996l-1.616-1.519l.138 2.172Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackPk1"><use href="#flagpackPk0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}