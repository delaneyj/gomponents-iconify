package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeNotice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTVolumeNotice0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#555" stroke-linejoin="round" d="M24 6v36c-7 0-12.201-9.16-12.201-9.16H6a2 2 0 0 1-2-2V17.01a2 2 0 0 1 2-2h5.799S17 6 24 6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M32 15h0a11.91 11.91 0 0 1 1.684 1.859A12.07 12.07 0 0 1 36 24c0 2.654-.846 5.107-2.278 7.09A11.936 11.936 0 0 1 32 33"/><path stroke-linecap="round" d="M34.236 41.186C40.084 37.696 44 31.305 44 24c0-7.192-3.796-13.496-9.493-17.02"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTVolumeNotice0)"/>`),
		g.Group(children),
	)
}