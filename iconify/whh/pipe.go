package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pipe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M768 1024q-51 0-100-21t-94.5-60.5t-81-79.5t-76.5-95q-17-23-42-49.5T309 657t-88-58t-93-23H32q-13 0-22.5-9.5T0 544t9.5-22.5T32 512h96q37 0 55.5.5T236 515t54.5 7t50 14t51 22.5t45 34T480 640q25 34 79.5 81t80.5 47q64 0 64-64V576q0-26 18.5-45t45.5-19h192q26 0 45 19t19 45v192q0 106-75 181t-181 75zm128-768q0 36 9 83q1 5 18 40.5t25 68.5h-58q-9-42-26-64q-23-24-27.5-47.5T832 256q0-80 37.5-136T960 64q-38 15-51 64t-13 128zm-192-51q0 38 9 89q50 56 54 154h-63q0-64-9-76q-55-62-55-167q0-85 37.5-145T768 0q15 0 32 7q-42 18-69 72.5T704 205z"/>`),
		g.Group(children),
	)
}