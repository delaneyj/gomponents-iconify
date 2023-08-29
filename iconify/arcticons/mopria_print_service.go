package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MopriaPrintService(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="24.842" height="7.238" x="11.58" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.5" ry="1.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.653 28.418h4.916c.604 0 1.09.42 1.09.94v1.616c0 .52-.486.939-1.09.939h-4.916m-.005-5.349h4.518c.842 0 1.518-.678 1.518-1.52V14.258c0-.841-.676-1.52-1.518-1.52H8.836c-.841 0-1.519.679-1.519 1.52v10.788c0 .841.678 1.52 1.52 1.52h4.945m.002-1.76H11.56c-.602 0-1.087-.42-1.087-.94V21.78c0-.52.485-.938 1.087-.938h24.88c.602 0 1.087.418 1.087.938v2.086c0 .52-.485.94-1.087.94h-1.787m2.298-10.132h.013a2.068 2.068 0 1 1-2.07 2.07v-.001c0-1.139.919-2.063 2.057-2.069Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.847 22.863h18.741c.588 0 1.065.519 1.065 1.158v17.32c0 .64-.477 1.159-1.064 1.159H14.846c-.588 0-1.064-.519-1.064-1.158v-17.32c0-.64.476-1.159 1.064-1.159Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.783 31.913H8.474c-.641 0-1.158-.419-1.158-.939v-1.617c0-.52.517-.938 1.158-.938h5.309m2.483 9.031h15.688m-15.688 2.56h15.688"/><ellipse cx="19.505" cy="29.742" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.868" ry="1.529"/><ellipse cx="28.32" cy="26.988" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.84" ry="1.545"/><ellipse cx="22.338" cy="23.755" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.65" ry=".892"/><ellipse cx="27.737" cy="33.453" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="5.023" ry="1.971"/>`),
		g.Group(children),
	)
}