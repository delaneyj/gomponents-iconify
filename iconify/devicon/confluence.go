package devicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Confluence(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<defs><linearGradient id="deviconConfluence0" x1="26.791" x2="11.792" y1="28.467" y2="19.855" gradientTransform="scale(4)" gradientUnits="userSpaceOnUse"><stop offset="0" stop-color="#0052cc"/><stop offset=".918" stop-color="#2380fb"/><stop offset="1" stop-color="#2684ff"/></linearGradient><linearGradient id="deviconConfluence1" x1="5.209" x2="20.208" y1="2.523" y2="11.136" gradientTransform="scale(4)" gradientUnits="userSpaceOnUse"><stop offset="0" stop-color="#0052cc"/><stop offset=".918" stop-color="#2380fb"/><stop offset="1" stop-color="#2684ff"/></linearGradient></defs><path fill="url(#deviconConfluence0)" d="M19.492 86.227a249.047 249.047 0 0 0-3.047 4.933c-.867 1.45-.433 3.336 1.016 4.207l19.863 12.188c1.45.87 3.332.433 4.203-1.016a139.349 139.349 0 0 1 2.899-4.934c7.832-12.91 15.804-11.46 30.011-4.64l19.72 9.281c1.593.727 3.335 0 4.058-1.45l9.426-21.323c.722-1.453 0-3.336-1.454-4.063c-4.203-1.887-12.464-5.805-19.714-9.43c-26.82-12.914-49.586-12.043-66.98 16.247zm0 0"/><path fill="url(#deviconConfluence1)" d="M108.508 37.773a249.047 249.047 0 0 0 3.047-4.933c.87-1.45.433-3.336-1.016-4.207L90.676 16.445c-1.45-.87-3.332-.433-4.203 1.016a133.55 133.55 0 0 1-2.899 4.934c-7.832 12.91-15.804 11.46-30.011 4.64l-19.72-9.281c-1.593-.727-3.331 0-4.058 1.45l-9.422 21.323c-.726 1.453 0 3.34 1.45 4.063c4.203 1.887 12.468 5.805 19.714 9.43c26.825 12.77 49.586 12.042 66.98-16.247zm0 0"/>`),
		g.Group(children),
	)
}