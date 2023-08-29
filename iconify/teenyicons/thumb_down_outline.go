package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbDownOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M3.5 9.5H3v.146l.078.122L3.5 9.5Zm2.698 4.24l.421-.27l-.421.27Zm2.667-1.51l-.448.223l.448-.224ZM7.5 9.5V9a.5.5 0 0 0-.447.724L7.5 9.5Zm7-5h.5v-.167l-.1-.133l-.4.3Zm-2.4-3.2l.4-.3l-.4.3ZM8.282 14.231l.257.429l-.257-.429ZM1 10V0H0v10h1Zm2.078-.232l2.698 4.24l.843-.537l-2.697-4.24l-.844.537Zm6.234 2.237L7.947 9.276l-.894.448l1.364 2.729l.895-.447ZM7.5 10h5V9h-5v1ZM15 7.5v-3h-1v3h1Zm-.1-3.3L12.5 1l-.8.6l2.4 3.2l.8-.6ZM10.5 0h-5v1h5V0ZM3 2.5v7h1v-7H3Zm9.5 7.5A2.5 2.5 0 0 0 15 7.5h-1A1.5 1.5 0 0 1 12.5 9v1Zm-7-10A2.5 2.5 0 0 0 3 2.5h1A1.5 1.5 0 0 1 5.5 1V0Zm3.039 14.66a2.034 2.034 0 0 0 .773-2.655l-.895.448c.242.483.07 1.071-.393 1.35l.515.857ZM12.5 1a2.5 2.5 0 0 0-2-1v1a1.5 1.5 0 0 1 1.2.6l.8-.6ZM5.776 14.008a2.034 2.034 0 0 0 2.763.652l-.515-.858a1.034 1.034 0 0 1-1.405-.331l-.843.537Z"/>`),
		g.Group(children),
	)
}