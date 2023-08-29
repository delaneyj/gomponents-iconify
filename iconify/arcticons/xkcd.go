package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xkcd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.657 16.243c1.846-.527 6.326-5.618 3.997-9.152s-8.925-3.59-10.28.675s-.025 7.96 2.193 8.523a6.038 6.038 0 0 0 4.09-.046Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.561 28.162c.417-.014 3.096-11.92 3.096-11.92s2.73 11.088 2.534 12.328"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m32.657 16.243l-.34 10.427s-3.083 8.36-3.083 9.245s-1.024 6.658-.964 7.181m4.047-16.425s3.08 11.806 3.383 12.138s1.596 4.288 1.851 4.303M14.093 18.695s-.514 1.542-1.763.285s-.906-1.755-.906-1.755s-.905.019-.972-1.136s.868-1.029.868-1.029s-.53-4.195 2.015-5.785s8.214-1.683 9.626 2.2s-.135 7.713-3.036 8.744s-4.706.03-5.831-1.524Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.922 11.448c-.162.266-1.145 3.17-7.481 3.114s-1.349 4.134-1.349 4.134m.53.677s-2.857 6.836-3.299 11.36m3.299-11.36s.845 8.247 2.279 11.378"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.622 19.373s-.757 7.585-.289 9.824a30.442 30.442 0 0 1 2.891 8.924c.667 4.119 1.412 5.379 1.412 5.379"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.333 29.197s-.824 5.665-1.65 8.738a28.632 28.632 0 0 0-.939 5.05"/>`),
		g.Group(children),
	)
}