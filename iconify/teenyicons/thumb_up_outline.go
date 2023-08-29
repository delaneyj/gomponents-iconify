package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbUpOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m3.5 5.5l-.422-.268L3 5.354V5.5h.5Zm2.698-4.24l.421.27l-.421-.27Zm2.667 1.51l-.448-.223l.448.224ZM7.5 5.5l-.447-.224A.5.5 0 0 0 7.5 6v-.5Zm7 5l.4.3l.1-.133V10.5h-.5Zm-2.4 3.2l.4.3l-.4-.3ZM8.282.769L8.539.34l-.257.43ZM0 5v10h1V5H0Zm3.922.768L6.619 1.53L5.776.992l-2.698 4.24l.844.536Zm4.495-3.22L7.053 5.275l.894.448l1.365-2.73l-.895-.447ZM7.5 6h5V5h-5v1ZM14 7.5v3h1v-3h-1Zm.1 2.7l-2.4 3.2l.8.6l2.4-3.2l-.8-.6ZM10.5 14h-5v1h5v-1ZM4 12.5v-7H3v7h1ZM12.5 6A1.5 1.5 0 0 1 14 7.5h1A2.5 2.5 0 0 0 12.5 5v1Zm-7 8A1.5 1.5 0 0 1 4 12.5H3A2.5 2.5 0 0 0 5.5 15v-1ZM8.024 1.198c.464.278.635.866.393 1.35l.895.446A2.034 2.034 0 0 0 8.539.34l-.515.858ZM11.7 13.4a1.5 1.5 0 0 1-1.2.6v1a2.5 2.5 0 0 0 2-1l-.8-.6ZM6.62 1.53c.3-.474.924-.62 1.404-.332L8.54.34a2.034 2.034 0 0 0-2.763.652l.843.537Z"/>`),
		g.Group(children),
	)
}