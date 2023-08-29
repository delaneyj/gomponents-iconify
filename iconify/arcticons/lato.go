package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lato(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2ZM5.255 20.694l36.863 20.975"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5.5 11.981l14.002 7.911l.952-1.632s2.397-1.417 5.78-.907c.501.076 2.176.884 2.176.884s3.377-4.363 5.01-1.473s-.477 4.675-1.803 4.692s-1.615-1.632-1.615-1.632"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.021 22.867c2.431-.544 5.32-2.173 5.32-2.173m-5.739 11.008c-.55-1.22.215-2.783.98-3.837s2.244-2.448-.56-4.998m-7.245 4.952c-.813-.413-.184-2.045.938-2.232s4.504-.442 4.436-1.394s-.442-1.14-.646-2.04s-.238-2.499 2.737-2.771m4.347 13.362c-.975.331-1.956.6-2.927.13m-13.121-8.789c.045.907.165 1.418 1.11 1.956"/>`),
		g.Group(children),
	)
}