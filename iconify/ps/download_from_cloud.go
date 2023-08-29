package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownloadFromCloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M427 171v-11q0-65-47.5-112.5T267 0q-57 0-101 35.5T111 126q-48 6-79.5 42.5T0 254q0 53 38 91.5t90 38.5h299q37-8 61-38t24-69q0-38-24-68.5T427 171zm0 170H128q-35 0-60-26t-25-61q0-31 25.5-54.5T128 171q4 0 11-11l10-32q8-39 41.5-62T267 43q48 0 82.5 34.5T384 160v32q0 9 6.5 15t14.5 6h13q22 5 36.5 23t14.5 41t-11 41.5t-31 22.5zM307 218l-30 21V128q0-21-21-21t-21 21v109l-30-22q-20-11-30 7q-10 19 6 30l64 42q2 0 2 3h22q2 0 2-3l64-42q18-11 6-30q-17-15-34-4z"/>`),
		g.Group(children),
	)
}