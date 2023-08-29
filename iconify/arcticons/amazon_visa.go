package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AmazonVisa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="24.536" height="14" x="11.753" y="10.668" rx="2.409" ry="2.409"/><path d="M14.003 20.84h3.568m2.684-.02h7.492m-15.914-6.285H36.17"/></g><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M32.28 29.7c1.113-.45 3.092-1.048 3.688-.326c.644.781-.17 2.477-.92 3.794"/><path d="M11.798 30.223c1.759 1.397 6.954 3.535 12.488 3.535c5.276 0 8.496-1.912 10.167-3.08"/></g>`),
		g.Group(children),
	)
}