package devicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pytest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#696969" d="M31.512 30.398h61.304a3.006 3.006 0 0 1 0 6.012H31.512a3.007 3.007 0 0 1-3.004-3.004a3.008 3.008 0 0 1 3.004-3.008zm0 0"/><path fill="#009fe3" d="M32.047 24.32H44.37v2.844H32.047zm0 0"/><path fill="#c7d302" d="M48.168 24.32h12.324v2.844H48.168zm0 0"/><path fill="#f07e16" d="M64.07 24.32h12.328v2.844H64.07zm0 0"/><path fill="#df2815" d="M79.91 24.32h12.324v2.844H79.91zm0 15.22h12.324v20.835H79.91zm0 0"/><path fill="#f07e16" d="M64.07 39.54h12.352v33.847H64.07zm0 0"/><path fill="#c7d302" d="M48.168 39.54h12.324v50.698H48.168zm0 0"/><path fill="#009fe3" d="M32.047 39.54H44.37v61.792H32.047zm0 0"/>`),
		g.Group(children),
	)
}