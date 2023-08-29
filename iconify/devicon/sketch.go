package devicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sketch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#fdb300" d="M29.62 12.85L64 9.24l34.38 3.61L125 48.29l-61 70.47L3 48.29z"/><path fill="#ea6c00" d="M27.71 48.29L64 118.76L3 48.29h24.71zm72.58 0L64 118.76l61-70.47h-24.71z"/><path fill="#fdad00" d="M27.71 48.29h72.58L64 118.76L27.71 48.29z"/><path fill="#fdd231" d="m64 9.24l-34.38 3.61l-1.91 35.44L64 9.24zm0 0l34.38 3.61l1.91 35.44L64 9.24z"/><path fill="#fdad00" d="M125 48.29L98.38 12.85l1.91 35.44H125z"/><path fill="#feeeb7" d="M64 9.24L27.71 48.29h72.58L64 9.24z"/><path fill="#fdad00" d="m3 48.29l26.62-35.44l-1.91 35.44H3z"/>`),
		g.Group(children),
	)
}