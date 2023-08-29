package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackSh0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackSh1)"><use href="#flagpackSh0"/><path fill="#2E42A5" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><g clip-path="url(#flagpackSh2)"><path fill="#2E42A5" d="M0 0h18v14H0z"/><path fill="#F7FCFF" d="m-2.004 13l3.96 1.737L18.09 1.889l2.09-2.582l-4.236-.58l-6.58 5.536l-5.297 3.73L-2.004 13Z"/><path fill="#F50100" d="m-1.462 14.217l2.018 1.008L19.429-.933h-2.833l-18.058 15.15Z"/><path fill="#F7FCFF" d="m20.004 13l-3.96 1.737L-.09 1.889L-2.18-.693l4.236-.58l6.58 5.536l5.297 3.73L20.004 13Z"/><path fill="#F50100" d="m19.87 13.873l-2.019 1.009l-8.036-6.918l-2.383-.773L-2.38-.684H.453l9.807 7.688l2.605.927l7.004 5.942Z"/><path fill="#F50100" fill-rule="evenodd" d="M9.985 0h-2v6H0v2h7.985v6h2V8H18V6H9.985V0Z" clip-rule="evenodd"/><path fill="#F7FCFF" fill-rule="evenodd" d="M6.485 0h5v4.5H18v5h-6.515L11.5 14h-5l-.015-4.5H0v-5h6.485V0Zm1.5 6H0v2h7.985v6h2V8H18V6H9.985V0h-2v6Z" clip-rule="evenodd"/></g></g><defs><clipPath id="flagpackSh1"><use href="#flagpackSh0"/></clipPath><clipPath id="flagpackSh2"><path fill="#fff" d="M0 0h18v14H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}