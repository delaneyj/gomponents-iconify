package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackHm0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackHm1)"><use href="#flagpackHm0"/><path fill="#2E42A5" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><g clip-path="url(#flagpackHm2)"><path fill="#2E42A5" d="M0 0h18v14H0z"/><path fill="#F7FCFF" d="m-2.004 13l3.96 1.737L18.09 1.889l2.09-2.582l-4.236-.58l-6.58 5.536l-5.297 3.73L-2.004 13Z"/><path fill="#F50100" d="m-1.462 14.217l2.018 1.008L19.429-.933h-2.833l-18.058 15.15Z"/><path fill="#F7FCFF" d="m20.004 13l-3.96 1.737L-.09 1.889L-2.18-.693l4.236-.58l6.58 5.536l5.297 3.73L20.004 13Z"/><path fill="#F50100" d="m19.87 13.873l-2.019 1.009l-8.036-6.918l-2.383-.773L-2.38-.684H.453l9.807 7.688l2.605.927l7.004 5.942Z"/><path fill="#F50100" fill-rule="evenodd" d="M9.985 0h-2v6H0v2h7.985v6h2V8H18V6H9.985V0Z" clip-rule="evenodd"/><path fill="#F7FCFF" fill-rule="evenodd" d="M6.485 0h5v4.5H18v5h-6.515L11.5 14h-5l-.015-4.5H0v-5h6.485V0Zm1.5 6H0v2h7.985v6h2V8H18V6H9.985V0h-2v6Z" clip-rule="evenodd"/></g><g fill="#F7FCFF" clip-path="url(#flagpackHm3)"><path d="m9 20.5l-1.302 1.203l.13-1.768l-1.753-.267l1.463-1.002l-.883-1.537l1.694.52L9 16l.65 1.648l1.695-.518l-.883 1.536l1.463 1.002l-1.752.267l.129 1.768L9 20.5ZM22 13l-.868.802l.086-1.178l-1.168-.179l.975-.668l-.589-1.024l1.13.346L22 10l.434 1.099l1.13-.346l-.59 1.024l.976.668l-1.168.178l.086 1.179L22 13Zm2-8l-.868.802l.086-1.179l-1.168-.178l.975-.668l-.589-1.024l1.13.346L24 2l.434 1.099l1.13-.346l-.59 1.024l.976.668l-1.168.178l.086 1.179L24 5Zm5 4l-.868.802l.086-1.179l-1.168-.178l.975-.668l-.589-1.024l1.13.346L29 6l.434 1.099l1.13-.346l-.59 1.024l.976.668l-1.168.178l.086 1.179L29 9Zm-4 13l-.868.802l.086-1.178l-1.168-.179l.975-.668l-.589-1.024l1.13.346L25 19l.434 1.099l1.13-.346l-.59 1.024l.976.668l-1.168.178l.086 1.179L25 22Zm3.5-7.75l-.882.463l.169-.981l-.714-.695l.986-.144L28.5 12l.44.893l.987.143l-.714.696l.169.982l-.882-.464Z"/></g></g><defs><clipPath id="flagpackHm1"><use href="#flagpackHm0"/></clipPath><clipPath id="flagpackHm2"><path fill="#fff" d="M0 0h18v14H0z"/></clipPath><clipPath id="flagpackHm3"><path fill="#fff" d="M6 2h25v21H6z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}