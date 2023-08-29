package academicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Elsevier(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M68.408 418.612V111.51l-63.673-4.898v-24.49c34.285 1.47 58.775 2.45 73.959 2.45h180.735L264.327 64h25.469v103.347h-25.47l-9.795-55.837H131.102v139.102h91.102l6.857-41.143h24v108.245h-24l-6.857-40.653h-91.102v109.715c0 24.49 1.47 27.428 2.939 28.897c1.47 1.47 6.367 2.94 46.53 2.94h36.245c29.388 0 43.592-1.96 49.96-8.327c6.857-6.857 14.693-25.47 24.49-66.123h24.49l-9.797 101.878H101.714c-34.285 0-71.02.98-97.47 1.959v-24l63.674-4.898z"/>`),
		g.Group(children),
	)
}