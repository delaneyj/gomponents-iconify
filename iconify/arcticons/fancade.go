package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fancade(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 43.5V20.958m0 0l15.491-7.874M21.54 6.461L24 5.21M8.509 13.084l5.5-2.795M24 36.595l15.491-7.874M24 20.958L8.509 13.084M24 36.595l-2.652-1.347m18.143-22.164v15.637m-9.867 1.932l-3.036 1.543v-3.577l3.036-1.543v3.577zm7.564-3.845l-3.036 1.543v-3.577l3.036-1.543v3.577z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.509 13.084V43.5h28.679V29.892M24 5.21l15.491 7.874M21.348 35.248V43.5m-5.381-20.633l2.746-1.396m-2.746 10.866l2.746-1.396m-2.746-8.074l-7.458-3.791m7.458 13.261l-7.458-3.79m7.458-5.68v9.47"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m8.509 19.076l2.747-1.396l7.457 3.791v9.47m8.056-14.166L24 15.367m7.745-1.122l-4.976 2.53v-3.15l4.976-2.53v3.15zm-4.976-.62L24 12.217m7.745-1.122l-2.768-1.407m0 0L24 12.217m0 3.15v-3.15m-2.334 1.964l-2.769-1.407m5.103.22l-2.334 1.187v-3.15l4.976-2.53v2.373m-4.976.157l-2.769-1.408m7.745-1.122l-2.768-1.407m0 0l-4.977 2.529m0 3.151V9.623m-2.334 1.964l-2.769-1.407m5.103.221l-2.334 1.186v-3.15l4.977-2.53V8.28m-4.977.157L13.794 7.03m7.746-1.123L18.771 4.5m0 0l-4.977 2.53m0 3.15V7.03M11.46 8.993l2.334 1.187m.018-2.382L11.46 8.993m0 0v2.556"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m14.179 27.441l-2.367-1.203v-3.007l2.367 1.204m0 0v3.006m20.255 7.196l2.754-1.4m-2.754 1.4v-3.343m0 3.343l-3.683-1.473m-4.058 5.426l2.754-1.4v-3.381m-2.754 4.781v-3.364m0 3.364L24 37.513"/>`),
		g.Group(children),
	)
}