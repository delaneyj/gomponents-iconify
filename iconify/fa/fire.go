package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fire(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1408 1696v64q0 13-9.5 22.5t-22.5 9.5H32q-13 0-22.5-9.5T0 1760v-64q0-13 9.5-22.5T32 1664h1344q13 0 22.5 9.5t9.5 22.5zM1152 640q0 78-24.5 144t-64 112.5t-87.5 88t-96 77.5t-87.5 72t-64 81.5T704 1312q0 96 67 224l-4-1l1 1q-90-41-160-83t-138.5-100T356 1230.5T283.5 1080T256 896q0-78 24.5-144t64-112.5t87.5-88t96-77.5t87.5-72t64-81.5T704 224q0-94-66-224l3 1l-1-1q90 41 160 83t138.5 100T1052 305.5t72.5 150.5t27.5 184z"/>`),
		g.Group(children),
	)
}