package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackSv0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackSv2)"><use href="#flagpackSv0"/><path fill="#F7FCFF" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackSv1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g mask="url(#flagpackSv1)"><path fill="#3D58DB" fill-rule="evenodd" d="M0 0v8h32V0H0Zm0 16v8h32v-8H0Z" clip-rule="evenodd"/><path stroke="#E8AA00" d="M19.425 11.858a3.642 3.642 0 1 1-7.283 0a3.642 3.642 0 0 1 7.283 0Z"/><path fill="#1E601B" fill-rule="evenodd" d="M13.81 9.662s-.952 1.568-.952 2.644s1.14 2.429 2.934 2.429c1.75 0 3.008-1.045 3.046-2.429c.038-1.384-.942-2.369-.942-2.369s.555 1.993.278 2.8c-.278.807-1.175 1.784-2.382 1.66c-1.207-.125-2.353-1.61-2.353-2.09c0-.48.372-2.645.372-2.645Z" clip-rule="evenodd"/><path stroke="#188396" d="M14.16 12.328h3.208"/><path stroke="#E8AA00" d="M14.459 11.806h2.76m-.169.605h-2.457l1.24-2.056l1.217 2.056Z"/></g></g><defs><clipPath id="flagpackSv2"><use href="#flagpackSv0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}