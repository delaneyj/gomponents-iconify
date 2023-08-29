package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RefreshOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTRefreshOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><rect width="10" height="10" x="17" y="24.071" fill="#555" stroke-linejoin="round" rx="2" transform="rotate(-45 17 24.071)"/><path d="M40.12 16c-2.945-5.927-9.06-10-16.129-10C16.924 6 10.945 10.073 8 16m0-8v8m6.78 0H8m0 16c2.945 5.927 9.061 10 16.129 10c7.067 0 13.046-4.073 15.991-10m0 8v-8m-6.78 0h6.78"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTRefreshOne0)"/>`),
		g.Group(children),
	)
}