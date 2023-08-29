package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GbUkm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackGbUkm0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackGbUkm3)"><use href="#flagpackGbUkm0"/><path fill="#2E42A5" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackGbUkm1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g mask="url(#flagpackGbUkm1)"><path fill="#fff" d="m-3.563 22.285l7.042 2.979l28.68-22.026l3.715-4.426l-7.53-.995l-11.698 9.491l-9.416 6.396l-10.793 8.581Z"/><path fill="#F50100" d="M-2.6 24.372L.989 26.1L34.54-1.599h-5.037l-32.102 25.97Z"/><path fill="#fff" d="m35.563 22.285l-7.042 2.979L-.159 3.238l-3.715-4.426l7.53-.995l11.698 9.491l9.416 6.396l10.793 8.581Z"/><path fill="#F50100" d="m35.323 23.783l-3.588 1.728l-14.286-11.86l-4.236-1.324l-17.445-13.5H.806l17.434 13.18l4.631 1.588l12.452 10.188Z"/><mask id="flagpackGbUkm2" fill="#fff"><path fill-rule="evenodd" d="M19.778-2h-7.556V8H-1.972v8h14.194v10h7.556V16h14.25V8h-14.25V-2Z" clip-rule="evenodd"/></mask><path fill="#F50100" fill-rule="evenodd" d="M19.778-2h-7.556V8H-1.972v8h14.194v10h7.556V16h14.25V8h-14.25V-2Z" clip-rule="evenodd"/><path fill="#fff" d="M12.222-2v-2h-2v2h2Zm7.556 0h2v-2h-2v2ZM12.222 8v2h2V8h-2ZM-1.972 8V6h-2v2h2Zm0 8h-2v2h2v-2Zm14.194 0h2v-2h-2v2Zm0 10h-2v2h2v-2Zm7.556 0v2h2v-2h-2Zm0-10v-2h-2v2h2Zm14.25 0v2h2v-2h-2Zm0-8h2V6h-2v2Zm-14.25 0h-2v2h2V8Zm-7.556-8h7.556v-4h-7.556v4Zm2 8V-2h-4V8h4Zm-16.194 2h14.194V6H-1.972v4Zm2 6V8h-4v8h4Zm12.194-2H-1.972v4h14.194v-4Zm2 12V16h-4v10h4Zm5.556-2h-7.556v4h7.556v-4Zm-2-8v10h4V16h-4Zm16.25-2h-14.25v4h14.25v-4Zm-2-6v8h4V8h-4Zm-12.25 2h14.25V6h-14.25v4Zm-2-12V8h4V-2h-4Z" mask="url(#flagpackGbUkm2)"/></g></g><defs><clipPath id="flagpackGbUkm3"><use href="#flagpackGbUkm0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}