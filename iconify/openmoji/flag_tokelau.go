package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagTokelau(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#1e50a0" d="M5 17h62v38H5z"/><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="m19.75 21.972l1.236 4L17.75 23.5h4l-3.236 2.472l1.236-4zm0 14l1.236 4L17.75 37.5h4l-3.236 2.472l1.236-4zm-7.25-8l1.236 4L10.5 29.5h4l-3.236 2.472l1.236-4zm14-1l.927 3L25 28.118h3l-2.427 1.854l.927-3z"/><path fill="#f1b31c" stroke="#f1b31c" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M56.496 22.004L22.5 47h34c-2.171-3.7-5.409-7.964-5.409-12.502c0-4.535 3.236-8.795 5.405-12.494Z"/><path fill="none" stroke="#f1b31c" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22.5 50h36"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}