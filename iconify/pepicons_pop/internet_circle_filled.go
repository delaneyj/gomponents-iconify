package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InternetCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopInternetCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M4.25 13a8.75 8.75 0 1 0 17.5 0a8.75 8.75 0 0 0-17.5 0Zm16 0a7.25 7.25 0 1 1-14.5 0a7.25 7.25 0 0 1 14.5 0Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M9.25 13c0 4.522 1.491 8.25 3.75 8.25s3.75-3.728 3.75-8.25S15.259 4.75 13 4.75S9.25 8.478 9.25 13Zm6 0c0 3.762-1.195 6.75-2.25 6.75s-2.25-2.988-2.25-6.75S11.945 6.25 13 6.25s2.25 2.988 2.25 6.75Z" clip-rule="evenodd"/><path d="m6.602 8.467l1.006-1.112c.1.09.209.18.325.267c1.271.952 3.3 1.54 5.515 1.54c1.891 0 3.652-.427 4.931-1.158c.308-.176.582-.367.819-.57l.974 1.141a6.73 6.73 0 0 1-1.048.73c-1.516.868-3.534 1.356-5.676 1.356c-2.522 0-4.865-.678-6.415-1.839a6.063 6.063 0 0 1-.431-.355Zm0 9.082l1.006 1.112c.1-.091.209-.18.325-.267c1.271-.952 3.3-1.54 5.515-1.54c1.891 0 3.652.427 4.931 1.158c.308.176.582.367.819.57l.974-1.141a6.841 6.841 0 0 0-1.048-.73c-1.516-.868-3.534-1.356-5.676-1.356c-2.522 0-4.865.678-6.415 1.839a6.06 6.06 0 0 0-.431.355ZM4.75 13.75v-1.5h16.5v1.5H4.75Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopInternetCircleFilled0)"/></g>`),
		g.Group(children),
	)
}