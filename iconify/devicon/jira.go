package devicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jira(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<defs><linearGradient id="deviconJira0" x1="22.034" x2="17.118" y1="9.773" y2="14.842" gradientTransform="scale(4)" gradientUnits="userSpaceOnUse"><stop offset=".176" stop-color="#0052cc"/><stop offset="1" stop-color="#2684ff"/></linearGradient><linearGradient id="deviconJira1" x1="16.641" x2="10.957" y1="15.564" y2="21.094" gradientTransform="scale(4)" gradientUnits="userSpaceOnUse"><stop offset=".176" stop-color="#0052cc"/><stop offset="1" stop-color="#2684ff"/></linearGradient></defs><path fill="#2684ff" d="M108.023 16H61.805c0 11.52 9.324 20.848 20.847 20.848h8.5v8.226c0 11.52 9.328 20.848 20.848 20.848V19.977A3.98 3.98 0 0 0 108.023 16zm0 0"/><path fill="url(#deviconJira0)" d="M85.121 39.04H38.902c0 11.519 9.325 20.847 20.844 20.847h8.504v8.226c0 11.52 9.328 20.848 20.848 20.848V43.016a3.983 3.983 0 0 0-3.977-3.977zm0 0"/><path fill="url(#deviconJira1)" d="M62.219 62.078H16c0 11.524 9.324 20.848 20.848 20.848h8.5v8.23c0 11.52 9.328 20.844 20.847 20.844V66.059a3.984 3.984 0 0 0-3.976-3.98zm0 0"/>`),
		g.Group(children),
	)
}