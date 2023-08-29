package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Youtobe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTYoutobe0"><g fill="#555" stroke="#fff" stroke-width="4"><path d="M44 32.768V15.232c0-1.325-.87-2.49-2.155-2.812C38.369 11.548 31.185 10 24 10c-7.184 0-14.369 1.548-17.845 2.42C4.87 12.743 4 13.907 4 15.232v17.536c0 1.325.87 2.49 2.155 2.812C9.631 36.452 16.815 38 24 38c7.184 0 14.369-1.548 17.845-2.42C43.13 35.257 44 34.093 44 32.768Z"/><path stroke-linejoin="round" d="M22.573 29.899a1 1 0 0 1-1.573-.82V18.921a1 1 0 0 1 1.573-.82l7.257 5.08a1 1 0 0 1 0 1.638l-7.256 5.08Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTYoutobe0)"/>`),
		g.Group(children),
	)
}