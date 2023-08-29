package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Skyward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 4.5l-3.607 2.082V9.37l-3.035 1.753v8.472l.626.362l-5.64 3.256V36.77L24 43.5M12.344 23.214L24 29.944m-6.016-9.988L24 23.43m-7.425-.216L24 27.5m-7.425-4.286l3.461-1.999m-2.678-10.093l1.391.803m2.048 1.183L24 14.957m-3.606-8.375L24 8.665m0 9.593l-1.605-.927m1.605.927l-1.605.927m-1.605-.927l1.605-.927m-1.605.927l1.605.927"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.395 19.185v2.1L24 20.36m-3.21-.001l1.605.927m-1.605-.927v-2.101m1.152 4.152v3.766m-1.193 5.696l-4.122-2.38l-1.57.905l4.086 2.36v4.764m.001-4.765l1.605-.884m-1.605 5.65l1.605-.885v-4.765"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m15.383 32.406l2.101 1.212v2.397m-1.654-1.338l-1.373-.792v1.519m0 0l1.373.792v-1.518M24 4.5l3.607 2.082V9.37l3.035 1.753v3.706m0 2.771v1.995l-.626.362l5.64 3.256V36.77L24 43.5m11.656-20.287L24 29.944m6.016-9.988L24 23.43m7.425-.217L24 27.5m7.425-4.287l-3.461-1.998m2.678-10.092L24 14.957m3.607-8.375L24 8.665m3.89 6.419l1.606-.927m-1.606.927l1.606.926m1.604-.926l-1.604-.927m1.604.927l-1.604.926m-.001 0v2.102l-1.605-.927m3.211 0l-1.605.927m1.605-.927v-2.101m-5.043 7.326v3.766"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m27.251 31.872l4.122-2.38l1.57.905l-4.086 2.36v4.764m0-4.764l-1.606-.885"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m28.857 37.521l-1.606-.885v-4.765m5.366.535l-2.1 1.213v2.396m1.653-1.337l1.373-.793v1.519"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m33.543 35.404l-1.373.793v-1.52m-10.172-24.38l-1.605-.927m1.605.927l-1.605.926m-1.605-.926l1.605-.927m-1.605.927l1.605.926m.001.001v2.1l1.605-.926m-3.211 0l1.605.927m-1.605-.927v-2.101m9.102 6.888v-2.101m-.283-5.714v1.754L24 13.206l-1.71-.987m-.292-1.923v2.101M24 14.957v3.3m0 2.103v3.07m0 0v4.07m0 2.443V43.5m0-30.293V8.665"/>`),
		g.Group(children),
	)
}