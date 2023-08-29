package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlobeAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3zM361 131q-32-56-92-76q19 35 29 76h63zM213 46q-27 39-40 85h81q-14-46-41-85zM48 259h72q-3-25-3-43t3-43H48q-5 23-5 43t5 43zm18 42q32 56 92 76q-19-35-29-76H66zm63-170q10-41 29-76q-60 20-92 76h63zm84 255q27-39 41-85h-81q13 46 40 85zm50-127q4-25 4-43t-4-43H163q-3 25-3 43t3 43h100zm6 118q60-20 92-76h-63q-10 41-29 76zm37-118h72q6-23 6-43t-6-43h-72q3 25 3 43t-3 43z"/>`),
		g.Group(children),
	)
}