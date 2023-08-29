package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RedditLoop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<mask id="lineMdRedditLoop0"><g fill="#fff" fill-opacity="0"><ellipse cx="12" cy="14.71" stroke="#fff" stroke-dasharray="48" stroke-dashoffset="48" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" rx="8" ry="5.29"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="48;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.4s" values="0;1"/></ellipse><circle cx="7.24" cy="11.97" r="2.24"><set attributeName="fill-opacity" begin="1s" to="1"/><animate fill="freeze" attributeName="cx" begin="1s" dur="0.2s" values="7.24;3.94"/></circle><circle cx="16.76" cy="11.97" r="2.24"><set attributeName="fill-opacity" begin="1s" to="1"/><animate fill="freeze" attributeName="cx" begin="1s" dur="0.2s" values="16.76;20.06"/></circle><circle cx="18.45" cy="4.23" r="1.61"><set attributeName="fill-opacity" begin="2.6s" to="1"/><animate attributeName="cx" begin="2.4s" dur="6s" repeatCount="indefinite" values="18.45;5.75;18.45"/></circle></g><path fill="none" stroke="#fff" stroke-dasharray="12" stroke-dashoffset="12" stroke-linecap="round" stroke-linejoin="round" stroke-width=".8" d="M12 8.75L13.18 3.11L18.21 4.18"><animate fill="freeze" attributeName="stroke-dashoffset" begin="2.4s" dur="0.2s" values="12;0"/><animate attributeName="d" begin="2.4s" dur="6s" repeatCount="indefinite" values="M12 8.75L13.18 3.11L18.21 4.18;M12 8.75L12 2L12 4.18;M12 8.75L10.82 3.11L5.79 4.18;M12 8.75L12 2L12 4.18;M12 8.75L13.18 3.11L18.21 4.18"/></path><g fill-opacity="0"><circle cx="8.45" cy="13.59" r="1.61"><animate fill="freeze" attributeName="fill-opacity" begin="1.2s" dur="0.4s" values="0;1"/></circle><circle cx="15.55" cy="13.59" r="1.61"><animate fill="freeze" attributeName="fill-opacity" begin="1.6s" dur="0.4s" values="0;1"/></circle></g><path fill="none" stroke="#000" stroke-dasharray="8" stroke-dashoffset="8" stroke-linecap="round" stroke-width=".8" d="M8.47 17.52C8.47 17.52 9.41 18.58 12 18.58C14.58 18.58 15.53 17.52 15.53 17.52"><animate fill="freeze" attributeName="stroke-dashoffset" begin="2s" dur="0.2s" values="8;0"/></path></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdRedditLoop0)"/>`),
		g.Group(children),
	)
}