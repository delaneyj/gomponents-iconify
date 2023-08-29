package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ModeHeatCoolOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.925 16.875Q5.2 16.225 4.1 14.712T3 11.25q0-1.975.938-3.513T6 5.15q1.125-1.05 2.063-1.6L9 3v2.475q0 .625.45 1.063t1.075.437q.35 0 .65-.15t.5-.425L12 6q.95.55 1.625 1.35t1.025 1.8l-1.675 1.675q-.05-.6-.288-1.175T12.05 8.6q-.35.2-.738.288t-.787.087q-1.1 0-1.988-.613T7.25 6.725q-.95.925-1.6 2.063T5 11.25q0 .775.275 1.463t.75 1.212q.05-.5.288-.938T6.9 12.2L9 10.1l2.15 2.1q.05.05.1.125t.1.125l-1.425 1.425q-.05-.075-.088-.125t-.087-.1L9 12.925l-.7.7q-.125.125-.213.288T8 14.275q0 .3.175.537t.45.363l-1.7 1.7ZM9 10.1Zm0 0ZM7.4 22L6 20.6L19.6 7L21 8.4L17.4 12H21v2h-5.6l-.5.5l1.5 1.5H21v2h-2.6l2.1 2.1l-1.4 1.4l-2.1-2.1V22h-2v-4.6l-1.5-1.5l-.5.5V22h-2v-3.6L7.4 22Z"/>`),
		g.Group(children),
	)
}