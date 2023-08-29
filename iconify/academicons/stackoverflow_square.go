package academicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackoverflowSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M48 32C21.5 32 0 53.5 0 80v352c0 26.5 21.5 48 48 48h352c26.5 0 48-21.5 48-48V80c0-26.5-21.5-48-48-48H48zm226.068 64l84.99 114.55l-22.908 16.997l-84.988-114.55L274.068 96zm-52.836 45.822l109.745 91.268l-18.11 21.803l-109.742-91.276l18.107-21.795zm-42.127 59.121l129.333 60.596l-11.823 25.867l-129.33-60.228l11.82-26.235zM154.72 265.61l139.674 29.186l-5.913 28.09L148.81 293.69l5.91-28.082zm-65.78 36.21h28.46v85.732h199.906v-85.733h28.449V416H88.94V301.818zm56.909 28.456h142.63v28.45H145.85v-28.45z"/>`),
		g.Group(children),
	)
}