package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserClockOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10 3.25a4.25 4.25 0 1 0 0 8.5a4.25 4.25 0 0 0 0-8.5ZM7.25 7.5a2.75 2.75 0 1 1 5.5 0a2.75 2.75 0 0 1-5.5 0Z" clip-rule="evenodd"/><path fill="currentColor" d="M3.75 17A2.25 2.25 0 0 1 6 14.75h.34c.027 0 .053.004.078.012l.866.283c.68.222 1.38.358 2.084.41a.302.302 0 0 0 .312-.239c.076-.327.174-.645.294-.952a.214.214 0 0 0-.192-.29a7.251 7.251 0 0 1-2.032-.355l-.866-.283a1.752 1.752 0 0 0-.543-.086H6A3.75 3.75 0 0 0 2.25 17v1.188c0 .754.546 1.396 1.29 1.517c2.157.353 4.337.527 6.517.524c.152 0 .247-.165.18-.3a6.958 6.958 0 0 1-.38-.914a.418.418 0 0 0-.388-.29a38.603 38.603 0 0 1-5.688-.5a.037.037 0 0 1-.031-.037V17Zm13.5-1.7a.75.75 0 1 0-1.5 0v1.773c0 .24.115.465.309.606l1 .728a.75.75 0 1 0 .882-1.214l-.691-.502V15.3Z"/><path fill="currentColor" fill-rule="evenodd" d="M16.5 22.3a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11Zm0-1.5a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}