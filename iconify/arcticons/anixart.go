package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Anixart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="26.125" r="7.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.006 8.767c-8.065.893-14.54 7.419-15.377 15.49c-.802 7.716 3.429 14.536 9.806 17.606a.818.818 0 0 0 1.168-.611l.74-4.975a.809.809 0 0 1 1.154-.607c1.366.649 2.89 1.017 4.503 1.017s3.138-.368 4.503-1.017a.809.809 0 0 1 1.154.607l.74 4.974a.818.818 0 0 0 1.166.613c5.858-2.82 9.906-8.803 9.906-15.74c0-10.304-8.923-18.525-19.464-17.358Zm-8.834.911c.59-.344.485-1.232-.16-1.456C10.048 7.195 5.155 5.614 4.6 6.17c-.502.502.994 5.236 2.007 8.227c.222.654 1.088.737 1.463.158a15.242 15.242 0 0 1 5.102-4.876Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.828 9.678c-.59-.344-.485-1.232.16-1.456c2.964-1.027 7.857-2.608 8.412-2.053c.502.502-.994 5.236-2.007 8.227c-.222.654-1.088.737-1.463.158a15.242 15.242 0 0 0-5.102-4.876Z"/>`),
		g.Group(children),
	)
}