package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngryBirdsSpace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.606 8.848c-1.163-.29-2.907-.775-5.135-.872c-1.453-.097-3.681-.194-3.778.485s2.131 1.259 4.747 3.487c1.744 1.26 2.422 2.132 3.1 3.1m7.654 21.702c-1.55.097-5.135-.775-5.716-1.55c-.775-.776 3.584-4.554 5.716-4.845"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.532 20.861c-.775-1.162-.775-.872-6.588-6.103c-6.394-5.716-6.297-6.007-8.041-7.266C17.896 3.52 11.115 3.035 10.921 3.81c-.194.678 4.553 1.841 7.944 6.491c1.163 1.744 1.744 2.52 2.035 3.682M7.918 24.543c-.872 2.228-2.035 5.231-1.26 9.106c1.163 5.426 5.813 8.138 6.879 8.816C17.025 44.5 20.222 44.5 24.291 44.5c4.166 0 8.235 0 12.013-2.906c2.616-2.035 3.681-4.65 4.166-5.523c1.84-4.456.968-8.816-.097-11.625M15.862 14.177c-1.26.871-3.875 2.809-6.103 6.49"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m23.128 30.646l.29-2.325l-6.006-1.647c-.969 2.131-.097 4.553 2.616 5.425m4.166 4.651c1.55-.097 4.553-.775 5.425-1.55c.775-.776-3.294-4.554-5.425-4.845"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m25.26 30.646l-.194-2.325l6.006-1.55c.969 2.131-.29 4.65-3.1 5.619m.581 3.585c.679 1.55-2.906 6.2-4.36 6.2c-.968 0-4.65-4.457-4.359-6.2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.65 33.262c6.2-.872 11.141-6.104 13.273-13.37L27.1 23.962c-.678.774-1.55 1.55-2.906 1.55c-1.453 0-2.713-.97-3.391-1.744l-14.726-4.07c2.035 7.267 7.072 12.498 13.273 13.37m-.969 1.841c-3.1.872-5.522 3.39-6.007 6.588m23.058.484c-.29-3.294-2.616-6.007-5.716-7.072"/>`),
		g.Group(children),
	)
}