package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackKn0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackKn1)"><use href="#flagpackKn0"/><path fill="#C51918" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><path fill="#5EAA22" fill-rule="evenodd" d="M0 0v24L32 0H0Z" clip-rule="evenodd"/><path fill="#272727" stroke="#FFD018" stroke-width="2.5" d="m.636 27.952l.723.807l.902-.602L38.52 3.954l1.2-.8l-.963-1.074l-5.32-5.936l-.724-.807l-.901.602l-36.261 24.203l-1.2.8l.963 1.074l5.32 5.936Z"/><path fill="#F7FCFF" fill-rule="evenodd" d="M12.089 16.918L11 18.813l-.695-2.134l-2.162-.688l1.801-1.1l-.217-2.233l1.778 1.367l1.842-1.02l-.516 2.273l1.477 1.701l-2.22-.06Zm10-6.553L21 12.26l-.695-2.134l-2.162-.688l1.801-1.1l-.217-2.233l1.778 1.367l1.842-1.02l-.516 2.273l1.477 1.701l-2.22-.06Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackKn1"><use href="#flagpackKn0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}