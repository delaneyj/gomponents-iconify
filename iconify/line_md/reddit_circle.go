package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RedditCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<mask id="lineMdRedditCircle0"><path fill="#fff" fill-opacity="0" stroke="#fff" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-width="2" d="M12 3C16.9706 3 21 7.02944 21 12C21 16.9706 16.9706 21 12 21C7.02944 21 3 16.9706 3 12C3 7.02944 7.02944 3 12 3Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="60;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.6s" dur="0.5s" values="0;1"/></path><g fill-opacity="0"><ellipse cx="12" cy="13.77" rx="5.83" ry="4.06"><animate fill="freeze" attributeName="fill-opacity" begin="1.1s" dur="0.4s" values="0;1"/></ellipse><circle cx="8.99" cy="11.99" r="1.45"><set attributeName="fill-opacity" begin="1.5s" to="1"/><animate fill="freeze" attributeName="cx" begin="1.5s" dur="0.2s" values="8.99;6.79"/></circle><circle cx="15.01" cy="11.99" r="1.45"><set attributeName="fill-opacity" begin="1.5s" to="1"/><animate fill="freeze" attributeName="cx" begin="1.5s" dur="0.2s" values="15.01;17.21"/></circle><circle cx="16.18" cy="7.01" r="1.04"><set attributeName="fill-opacity" begin="3.1s" to="1"/></circle></g><path fill="none" stroke="#000" stroke-dasharray="8" stroke-dashoffset="8" stroke-linecap="round" stroke-linejoin="round" stroke-width=".54" d="M12 9.91L12.76 6.27L16 6.98"><animate fill="freeze" attributeName="stroke-dashoffset" begin="2.9s" dur="0.2s" values="8;0"/></path><g fill="#fff" fill-opacity="0"><circle cx="9.71" cy="13.04" r="1.04"><animate fill="freeze" attributeName="fill-opacity" begin="1.7s" dur="0.4s" values="0;1"/></circle><circle cx="14.29" cy="13.04" r="1.04"><animate fill="freeze" attributeName="fill-opacity" begin="2.1s" dur="0.4s" values="0;1"/></circle></g><path fill="none" stroke="#fff" stroke-dasharray="6" stroke-dashoffset="6" stroke-linecap="round" stroke-width=".54" d="M9.72 15.6C9.72 15.6 10.33 16.29 12 16.291C13.67 16.29 14.28 15.6 14.28 15.6"><animate fill="freeze" attributeName="stroke-dashoffset" begin="2.5s" dur="0.2s" values="6;0"/></path></mask><rect width="24" height="24" fill="currentColor" mask="url(#lineMdRedditCircle0)"/>`),
		g.Group(children),
	)
}