package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChairAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M429 1056h676l72-748q14-147-41-227.5T960 0H576Q455 0 399 80.5T357 308zm595 288H512v252q0 28-18.5 48t-45.5 20H320q-27 0-45.5-18t-18.5-46l-1-271q-35-15-61.5-38.5t-43-57.5t-27-67.5t-15.5-83t-6.5-89.5t-1.5-100q-44-9-72-42.5T0 768q0-55 38.5-91.5T158 640q35 0 59.5 10t40.5 31.5t24 38t18 46.5l39 331q8 41 16 47q10 8 44 8h736q37 0 44.5-8t15.5-47l39-331q10-30 18-46.5t24-38t40.5-31.5t59.5-10q81 0 120.5 36.5T1536 768q0 48-28.5 82t-73.5 43q0 61-1.5 100.5t-6.5 90t-15 84.5t-26.5 68.5t-43 58t-61.5 38.5v267q0 28-18.5 46t-45.5 18h-128q-27 0-45.5-18t-18.5-46v-256z"/>`),
		g.Group(children),
	)
}