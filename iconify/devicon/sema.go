package devicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sema(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path d="M63.93 0C28.633 0 0 28.664 0 64s28.633 64 63.93 64c35.293 0 63.93-28.664 63.93-64S99.222 0 63.93 0ZM32.469 101.656l22.035-81.922h12.25l-22.031 81.922Zm70.355-30.094l-6.691 5.418l-10.645-12.085L74.395 76.98l-6.957-5.12l8.8-13.368L61.73 52.81l2.586-9.258l15.313 3.988l1.277-15.21h8.684l1.277 15.21l14.512-3.691l2.973 8.96l-14.36 5.684Zm0 0"/>`),
		g.Group(children),
	)
}