package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BroadcastTower(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M7.188 4.188c-4.297 4.183-4.282 11.125 0 15.406l1.406-1.407c-3.52-3.519-3.504-9.148 0-12.562zm17.625.093L23.405 5.72c3.524 3.523 3.524 9.039 0 12.562l1.407 1.438a10.897 10.897 0 0 0 0-15.438zM9.905 7.188c-2.586 2.585-2.586 6.82 0 9.406l1.406-1.407a4.678 4.678 0 0 1 0-6.593zm12.188.093L20.687 8.72a4.64 4.64 0 0 1 0 6.562l1.407 1.438c2.586-2.586 2.586-6.852 0-9.438zM16 10a2 2 0 0 0-2 2c0 .625.3 1.164.75 1.531L10.312 26H9v2h4v-2h-.594L16 15.969L19.594 26H19v2h4v-2h-1.313L17.25 13.531c.45-.367.75-.906.75-1.531a2 2 0 0 0-2-2z"/>`),
		g.Group(children),
	)
}