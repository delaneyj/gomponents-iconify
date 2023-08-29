package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AppmgrIii(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m9.098 21.15l-.84-7.583H7.85c-1.052 0-1.9.845-1.9 1.896v5.687m36.1 1.896v-7.583c0-1.05-.847-1.896-1.9-1.896h-7.129M5.951 27.785c0 .168.027.329.068.484l-.069-.484Zm18.456-7.991l-.553-4.185c-.054-.487-.458-.852-.947-.892a1.134 1.134 0 0 0-.106-.004h0a1.01 1.01 0 0 0-.109.006c-.581.064-1 .564-.938 1.12m.688 6.26l-.851-7.69m-8.253.909l.646 5.833m-.439-3.963a1.022 1.022 0 0 0-1.052-.896a.979.979 0 0 0-.108.007c-.582.064-1 .564-.94 1.12l.414 3.732"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m26.21 12.955l-.47-4.246c-.106-.957-.745-1.867-2.195-1.756L9.697 8.48a1.982 1.982 0 0 0-1.76 2.191l1.16 10.48"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m40.015 23.046l1.078-6.021a1.984 1.984 0 0 0-1.611-2.306l-11.767-2.097a2.017 2.017 0 0 0-.403-.031h-.001a1.987 1.987 0 0 0-1.908 1.638l-1.239 6.922"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.85 41.056L5 21.15h13.3l.95.948h3.8v-.948h4.75l2.85 1.896H43l-2.85 18.01H7.85Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m8.8 25.89l1.705 12.297H27.55l1.2-12.297H8.8Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.142 36.169V28.88h1.156c1.958 0 3.559 1.64 3.559 3.644s-1.601 3.644-3.559 3.644h-1.156Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.42 35.349c.445.546.98.82 1.78.82h1.067c.979 0 1.78-.82 1.78-1.822s-.801-1.822-1.78-1.822h-1.156c-.979 0-1.78-.82-1.78-1.822s.801-1.822 1.78-1.822h1.067c.801 0 1.335.182 1.78.82" class="e"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.348 25.89h6.757m-7.274 3.317h6.757m-7.274 3.318h6.757m-7.274 3.317h6.757M31.28 39.16h6.757M13.338 15.318l8.253-.91s-.83-3.986-4.563-3.495c-4.098.539-3.69 4.405-3.69 4.405Zm1.549-3.615l-2.06-1.526m6.442 1.053l1.639-1.958"/><circle cx="15.846" cy="13.34" r=".747" fill="currentColor"/><circle cx="18.522" cy="13.032" r=".747" fill="currentColor"/>`),
		g.Group(children),
	)
}