package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yelp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m11.5 11l-.458-8.24a.6.6 0 0 0-.771-.541L6.814 3.256a.6.6 0 0 0-.311.93L11.5 11Zm2.5 1.5l4.57-.83a.6.6 0 0 0 .38-.94l-1.445-2.023a.6.6 0 0 0-.987.016L14 12.5Zm.5 3.5l2.066 4.132a.6.6 0 0 0 1.017.091l1.835-2.446a.6.6 0 0 0-.373-.95L14.5 16Zm-3 .5l-3.341 3.341a.6.6 0 0 0 .213.986l2.317.869a.6.6 0 0 0 .811-.562V16.5Zm-2-2.5l-4.132-2.066a.6.6 0 0 0-.868.537v2.643a.6.6 0 0 0 .823.557L9.5 14Z"/>`),
		g.Group(children),
	)
}