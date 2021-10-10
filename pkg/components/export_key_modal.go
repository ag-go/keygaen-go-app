package components

import (
	app "github.com/maxence-charriere/go-app/v9/pkg/app"
)

const (
	exportPublicKeyForm  = "export-public-key-form"
	exportPrivateKeyForm = "export-private-key-form"
)

type ExportKeyModal struct {
	app.Compo

	OnDownloadPublicKey func(armor bool)
	OnViewPublicKey     func()

	OnDownloadPrivateKey func(armor bool)
	OnViewPrivateKey     func()

	OnOK func()

	skipPublicKeyArmor  bool
	skipPrivateKeyArmor bool
}

func (c *ExportKeyModal) Render() app.UI {
	return &Modal{
		ID:    "export-key-modal",
		Title: "Export Key",
		Body: []app.UI{
			app.Div().
				Class("pf-c-card pf-m-compact pf-m-flat").
				Body(
					app.Div().
						Class("pf-c-card__title").
						Body(
							app.I().
								Class("fas fa-globe pf-u-mr-sm"),
							app.Text("Public Key"),
						),
					app.Div().
						Class("pf-c-card__body").
						Body(
							app.P().
								Text("Anyone can use this key to encrypt messages to you and verify your identity; you may share it with the public."),
							app.Form().
								Class("pf-c-form pf-u-mt-lg").
								ID(exportPublicKeyForm).
								OnSubmit(func(ctx app.Context, e app.Event) {
									e.PreventDefault()
								}).
								Body(
									app.Div().
										Aria("role", "group").
										Class("pf-c-form__group").
										Body(
											app.Div().
												Class("pf-c-form__group-control").
												Body(
													app.Div().
														Class("pf-c-check").
														Body(
															&Controlled{
																Component: app.Input().
																	Class("pf-c-check__input").
																	Type("checkbox").
																	ID("armor-checkbox").
																	OnInput(func(ctx app.Context, e app.Event) {
																		c.skipPublicKeyArmor = !c.skipPublicKeyArmor
																	}),
																Properties: map[string]interface{}{
																	"checked": !c.skipPublicKeyArmor,
																},
															},
															app.Label().
																Class("pf-c-check__label").
																For("armor-checkbox").
																Body(
																	app.I().
																		Class("fas fa-shield-alt pf-u-mr-sm"),
																	app.Text("Armor"),
																),
															app.Span().
																Class("pf-c-check__description").
																Text("To increase portability, ASCII armor the key."),
														),
												),
										),
								),
						),
					app.Div().
						Class("pf-c-card__footer").
						Body(
							app.Button().
								Class("pf-c-button pf-m-control pf-u-mr-sm").
								Type("submit").
								Form(exportPublicKeyForm).
								OnClick(func(ctx app.Context, e app.Event) {
									c.OnDownloadPublicKey(!c.skipPublicKeyArmor)
								}).
								Body(
									app.Span().
										Class("pf-c-button__icon pf-m-start").
										Body(
											app.I().
												Class("fas fa-download").
												Aria("hidden", true),
										),
									app.Text("Download public key"),
								),
							app.If(
								!c.skipPublicKeyArmor,
								app.Button().
									Class("pf-c-button pf-m-control").
									Type("submit").
									Form(exportPublicKeyForm).
									OnClick(func(ctx app.Context, e app.Event) {
										c.OnViewPublicKey()
									}).
									Body(
										app.Span().
											Class("pf-c-button__icon pf-m-start").
											Body(
												app.I().
													Class("fas fa-eye").
													Aria("hidden", true),
											),
										app.Text("View public key"),
									),
							),
						),
				),
		},
		Footer: []app.UI{
			app.Button().
				Class("pf-c-button pf-m-primary").
				Type("button").
				Text("OK").
				OnClick(func(ctx app.Context, e app.Event) {
					c.clear()
					c.OnOK()
				}),
		},
		OnClose: func() {
			c.clear()
			c.OnOK()
		},
	}
}

func (c *ExportKeyModal) clear() {
	c.skipPublicKeyArmor = false
	c.skipPrivateKeyArmor = false
}