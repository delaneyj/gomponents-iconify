package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrueIservice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.686 7.84v4.72c0 1.27.755 2.48 1.963 2.874a2.936 2.936 0 0 0 3.886-2.769V7.84m0 4.971v2.925m-11.409-4.972A2.933 2.933 0 0 1 20.05 7.84m-2.924 0v7.896M12.589 5.5v8.774c0 .807.655 1.462 1.462 1.462h.439M10.98 7.84h3.217m22.384 6.434c-.439.877-1.462 1.462-2.486 1.462a2.933 2.933 0 0 1-2.924-2.925v-1.795c0-1.27.755-2.48 1.963-2.874a2.936 2.936 0 0 1 3.886 2.768v1.024h-5.85m-1.287 20.26c.076.1.146.207.21.32l3.568 6.292c.712 1.256.38 2.78-.743 3.417c-1.124.637-2.602.14-3.314-1.117l-3.568-6.291a2.939 2.939 0 0 1-.123-.242"/><ellipse cx="23.696" cy="26.383" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="5.87" ry="5.952" transform="rotate(-44.899 23.696 26.383)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.85 24.418c.79 1.268.782 2.96-.247 4.302a3.472 3.472 0 0 1-1.7 1.188c-2.564.81-4.922-1.082-4.922-3.525a3.716 3.716 0 0 1 4.555-3.62"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.851 24.415c-.161.694-.81 1.189-1.56 1.189c-.88 0-1.596-.678-1.596-1.515c0-.555.32-1.065.834-1.33"/><ellipse cx="23.696" cy="26.383" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="8.456" ry="8.524" transform="rotate(-44.823 23.696 26.383)"/>`),
		g.Group(children),
	)
}