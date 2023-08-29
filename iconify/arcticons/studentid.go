package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Studentid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 20.956l-15.885-8.18L24 4.5l15.885 8.276L24 20.956z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.246 17.437c-.095.57-.19 1.046-.19 1.617c0 3.9 3.14 6.944 6.944 6.944s6.944-3.14 6.944-6.944c0-.571-.095-1.047-.19-1.617"/><ellipse cx="24" cy="36.842" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="14.173" ry="6.659"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.115 12.776v9.131"/>`),
		g.Group(children),
	)
}