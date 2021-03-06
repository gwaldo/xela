package actions

import (
	"github.com/PagerDuty/xela/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Event)
// DB Table: Plural (events)
// Resource: Plural (Events)
// Path: Plural (/events)
// View Template Folder: Plural (/templates/events/)

// EventsResource is the resource for the Event model
type EventsResource struct {
	buffalo.Resource
}

// List gets all Events. This function is mapped to the path
// GET /events
func (v EventsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	events := &models.Events{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Events from the DB
	if err := q.All(events); err != nil {
		return errors.WithStack(err)
	}

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.Auto(c, events))
}

// Show gets the data for one Event. This function is mapped to
// the path GET /events/{event_id}
func (v EventsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Event
	event := &models.Event{}

	// To find the Event the parameter event_id is used.
	if err := tx.Find(event, c.Param("event_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, event))
}

// New renders the form for creating a new Event.
// This function is mapped to the path GET /events/new
func (v EventsResource) New(c buffalo.Context) error {
	return c.Render(200, r.Auto(c, &models.Event{}))
}

// Create adds a Event to the DB. This function is mapped to the
// path POST /events
func (v EventsResource) Create(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	user := c.Value("current_user").(*models.User)

	// Allocate an empty Event
	event := &models.Event{
		UserID:    user.ID,
		UpdatedBy: user.ID,
	}

	// Bind event to the html form elements
	if err := c.Bind(event); err != nil {
		return errors.WithStack(err)
	}

	event.LogoName = event.Logo.Filename

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(event)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)
		c.Flash().Add("danger", "Error with data. Maybe your dates are out of order?")

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, event))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Event was created successfully")

	// and redirect to the events index page
	return c.Render(201, r.Auto(c, event))
}

// Edit renders a edit form for a Event. This function is
// mapped to the path GET /events/{event_id}/edit
func (v EventsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Event
	event := &models.Event{}

	if err := tx.Find(event, c.Param("event_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, event))
}

// Update changes a Event in the DB. This function is mapped to
// the path PUT /events/{event_id}
func (v EventsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	user := c.Value("current_user").(*models.User)

	// Allocate an empty Event
	event := &models.Event{
		UpdatedBy: user.ID,
	}

	if err := tx.Find(event, c.Param("event_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind Event to the html form elements
	if err := c.Bind(event); err != nil {
		return errors.WithStack(err)
	}

	if event.Logo.Valid() {
		event.LogoName = event.Logo.Filename
	}
	verrs, err := tx.ValidateAndUpdate(event)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, event))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Event was updated successfully")

	// and redirect to the events index page
	return c.Render(200, r.Auto(c, event))
}

// Destroy deletes a Event from the DB. This function is mapped
// to the path DELETE /events/{event_id}
func (v EventsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Event
	event := &models.Event{}

	// To find the Event the parameter event_id is used.
	if err := tx.Find(event, c.Param("event_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(event); err != nil {
		return errors.WithStack(err)
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", "Event was destroyed successfully")

	// Redirect to the events index page
	return c.Render(200, r.Auto(c, event))
}
