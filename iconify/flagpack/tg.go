package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackTg0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackTg1)"><use href="#flagpackTg0"/><path fill="#5EAA22" stroke="#F7FCFF" stroke-width="2" d="M0-1h-1v26h34V-1H0Z"/><path fill="#FECA00" fill-rule="evenodd" d="M0 6v4h32V6H0Zm0 8v4h32v-4H0Z" clip-rule="evenodd"/><path fill="#F50101" d="M0 0h16v14H0z"/><path fill="#F7FCFF" fill-rule="evenodd" d="m8.25 10.144l-3.72 2.302l1.445-3.864L3 5.98h3.648L8.25 2.13l1.223 3.85h3.604l-2.532 2.603l1.246 3.674l-3.541-2.112Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackTg1"><use href="#flagpackTg0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}