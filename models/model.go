package models

import "time"

type Description struct {
	Format string `json:"format,omitempty"`
	Raw    string `json:"raw,omitempty"`
	HTML   string `json:"html,omitempty"`
}

type Elements struct {
	Elements []interface{} `json:"elements,omitempty"`
}

type Links struct {
	Href  string `json:"href,omitempty"`
	Title string `json:"title,omitempty"`
}

type Attachments struct {
	Type     string   `json:"_type,omitempty"`
	Total    int      `json:"total,omitempty"`
	Count    int      `json:"count,omitempty"`
	Embedded Elements `json:"_embedded,omitempty"`
	Links    struct {
		Self struct {
			Href string `json:"href,omitempty"`
		} `json:"self,omitempty"`
	} `json:"_links,omitempty"`
}

type Relations struct {
	Type     string `json:"_type,omitempty"`
	Total    int    `json:"total,omitempty"`
	Count    int    `json:"count,omitempty"`
	Embedded struct {
		Elements []interface{} `json:"elements,omitempty"`
	} `json:"_embedded,omitempty"`
	Links struct {
		Self struct {
			Href string `json:"href,omitempty"`
		} `json:"self,omitempty"`
	} `json:"_links,omitempty"`
}

type Type struct {
	Type        string    `json:"_type,omitempty"`
	ID          int       `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Color       string    `json:"color,omitempty"`
	Position    int       `json:"position,omitempty"`
	IsDefault   bool      `json:"isDefault,omitempty"`
	IsMilestone bool      `json:"isMilestone,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty"`
	Links       struct {
		Self struct {
			Href  string `json:"href,omitempty"`
			Title string `json:"title,omitempty"`
		} `json:"self,omitempty"`
	} `json:"_links,omitempty"`
}

type Priority struct {
	Type      string `json:"_type,omitempty"`
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Position  int    `json:"position,omitempty"`
	Color     string `json:"color,omitempty"`
	IsDefault bool   `json:"isDefault,omitempty"`
	IsActive  bool   `json:"isActive,omitempty"`
	Links     struct {
		Self struct {
			Href  string `json:"href,omitempty"`
			Title string `json:"title,omitempty"`
		} `json:"self,omitempty"`
	} `json:"_links,omitempty"`
}

type Project struct {
	Type        string `json:"_type,omitempty"`
	ID          int    `json:"id,omitempty"`
	Identifier  string `json:"identifier,omitempty"`
	Name        string `json:"name,omitempty"`
	Active      bool   `json:"active,omitempty"`
	Public      bool   `json:"public,omitempty"`
	Description struct {
		Format string `json:"format,omitempty"`
		Raw    string `json:"raw,omitempty"`
		HTML   string `json:"html,omitempty"`
	} `json:"description,omitempty"`
	CreatedAt         time.Time `json:"createdAt,omitempty"`
	UpdatedAt         time.Time `json:"updatedAt,omitempty"`
	Status            string    `json:"status,omitempty"`
	StatusExplanation struct {
		Format string `json:"format,omitempty"`
		Raw    string `json:"raw,omitempty"`
		HTML   string `json:"html,omitempty"`
	} `json:"statusExplanation,omitempty"`
	Links struct {
		Self struct {
			Href  string `json:"href,omitempty"`
			Title string `json:"title,omitempty"`
		} `json:"self,omitempty"`
		CreateWorkPackage struct {
			Href   string `json:"href,omitempty"`
			Method string `json:"method,omitempty"`
		} `json:"createWorkPackage,omitempty"`
		CreateWorkPackageImmediately struct {
			Href   string `json:"href,omitempty"`
			Method string `json:"method,omitempty"`
		} `json:"createWorkPackageImmediately,omitempty"`
		WorkPackages struct {
			Href string `json:"href,omitempty"`
		} `json:"workPackages,omitempty"`
		Categories struct {
			Href string `json:"href,omitempty"`
		} `json:"categories,omitempty"`
		Versions struct {
			Href string `json:"href,omitempty"`
		} `json:"versions,omitempty"`
		Memberships struct {
			Href string `json:"href,omitempty"`
		} `json:"memberships,omitempty"`
		Types struct {
			Href string `json:"href,omitempty"`
		} `json:"types,omitempty"`
		Update struct {
			Href   string `json:"href,omitempty"`
			Method string `json:"method,omitempty"`
		} `json:"update,omitempty"`
		UpdateImmediately struct {
			Href   string `json:"href,omitempty"`
			Method string `json:"method,omitempty"`
		} `json:"updateImmediately,omitempty"`
		Delete struct {
			Href   string `json:"href,omitempty"`
			Method string `json:"method,omitempty"`
		} `json:"delete,omitempty"`
		Schema struct {
			Href string `json:"href,omitempty"`
		} `json:"schema,omitempty"`
		Parent struct {
			Uncacheable bool        `json:"uncacheable,omitempty"`
			Href        interface{} `json:"href,omitempty"`
			Title       interface{} `json:"title,omitempty"`
		} `json:"parent,omitempty"`
	} `json:"_links,omitempty"`
}

type Status struct {
	Type             string      `json:"_type,omitempty"`
	ID               int         `json:"id,omitempty"`
	Name             string      `json:"name,omitempty"`
	IsClosed         bool        `json:"isClosed,omitempty"`
	Color            string      `json:"color,omitempty"`
	IsDefault        bool        `json:"isDefault,omitempty"`
	IsReadonly       bool        `json:"isReadonly,omitempty"`
	DefaultDoneRatio interface{} `json:"defaultDoneRatio,omitempty"`
	Position         int         `json:"position,omitempty"`
	Links            struct {
		Self struct {
			Href  string `json:"href,omitempty"`
			Title string `json:"title,omitempty"`
		} `json:"self,omitempty"`
	} `json:"_links,omitempty"`
}

type Author struct {
	Type        string      `json:"_type,omitempty"`
	ID          int         `json:"id,omitempty"`
	Name        string      `json:"name,omitempty"`
	CreatedAt   time.Time   `json:"createdAt,omitempty"`
	UpdatedAt   time.Time   `json:"updatedAt,omitempty"`
	Login       string      `json:"login,omitempty"`
	Admin       bool        `json:"admin,omitempty"`
	FirstName   string      `json:"firstName,omitempty"`
	LastName    string      `json:"lastName,omitempty"`
	Email       string      `json:"email,omitempty"`
	Avatar      string      `json:"avatar,omitempty"`
	Status      string      `json:"status,omitempty"`
	IdentityURL interface{} `json:"identityUrl,omitempty"`
	Links       struct {
		Self struct {
			Href  string `json:"href,omitempty"`
			Title string `json:"title,omitempty"`
		} `json:"self,omitempty"`
		Memberships struct {
			Href  string `json:"href,omitempty"`
			Title string `json:"title,omitempty"`
		} `json:"memberships,omitempty"`
		ShowUser struct {
			Href string `json:"href,omitempty"`
			Type string `json:"type,omitempty"`
		} `json:"showUser,omitempty"`
		UpdateImmediately struct {
			Href   string `json:"href,omitempty"`
			Title  string `json:"title,omitempty"`
			Method string `json:"method,omitempty"`
		} `json:"updateImmediately"`
		Lock struct {
			Href   string `json:"href"`
			Title  string `json:"title"`
			Method string `json:"method"`
		} `json:"lock"`
		Delete struct {
			Href   string `json:"href"`
			Title  string `json:"title"`
			Method string `json:"method"`
		} `json:"delete"`
	} `json:"_links"`
}

type Assignee struct {
	Type        string      `json:"_type"`
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
	Login       string      `json:"login"`
	Admin       bool        `json:"admin"`
	FirstName   string      `json:"firstName"`
	LastName    string      `json:"lastName"`
	Email       string      `json:"email"`
	Avatar      string      `json:"avatar"`
	Status      string      `json:"status"`
	IdentityURL interface{} `json:"identityUrl"`
	Links       struct {
		Self struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"self"`
		Memberships struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"memberships"`
		ShowUser struct {
			Href string `json:"href"`
			Type string `json:"type"`
		} `json:"showUser"`
		UpdateImmediately struct {
			Href   string `json:"href"`
			Title  string `json:"title"`
			Method string `json:"method"`
		} `json:"updateImmediately"`
		Lock struct {
			Href   string `json:"href"`
			Title  string `json:"title"`
			Method string `json:"method"`
		} `json:"lock"`
		Delete struct {
			Href   string `json:"href"`
			Title  string `json:"title"`
			Method string `json:"method"`
		} `json:"delete"`
	} `json:"_links"`
}

type Version struct {
	Type        string `json:"_type"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description struct {
		Format string `json:"format"`
		Raw    string `json:"raw"`
		HTML   string `json:"html"`
	} `json:"description"`
	StartDate interface{} `json:"startDate"`
	EndDate   interface{} `json:"endDate"`
	Status    string      `json:"status"`
	Sharing   string      `json:"sharing"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
	Links     struct {
		Self struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"self"`
		Schema struct {
			Href string `json:"href"`
		} `json:"schema"`
		Update struct {
			Href   string `json:"href"`
			Method string `json:"method"`
		} `json:"update"`
		UpdateImmediately struct {
			Href   string `json:"href"`
			Method string `json:"method"`
		} `json:"updateImmediately"`
		Delete struct {
			Href   string `json:"href"`
			Method string `json:"method"`
		} `json:"delete"`
		AvailableInProjects struct {
			Href string `json:"href"`
		} `json:"availableInProjects"`
	} `json:"_links"`
}

type Embedded struct {
	Attachments   Attachments   `json:"attachments"`
	Relations     Relations     `json:"relations"`
	Type          Type          `json:"type"`
	Priority      Priority      `json:"priority"`
	Project       Project       `json:"project"`
	Status        Status        `json:"status"`
	Author        Author        `json:"author"`
	Assignee      Assignee      `json:"assignee"`
	Version       Version       `json:"version"`
	CustomActions []interface{} `json:"customActions"`
}

type WorkPackage struct {
	Action      string `json:"action"`
	WorkPackage struct {
		Type                 string      `json:"_type"`
		ID                   int         `json:"id"`
		LockVersion          int         `json:"lockVersion"`
		Subject              string      `json:"subject"`
		Description          Description `json:"description"`
		StartDate            interface{} `json:"startDate"`
		DueDate              interface{} `json:"dueDate"`
		EstimatedTime        interface{} `json:"estimatedTime"`
		DerivedEstimatedTime interface{} `json:"derivedEstimatedTime"`
		PercentageDone       int         `json:"percentageDone"`
		CreatedAt            time.Time   `json:"createdAt"`
		UpdatedAt            time.Time   `json:"updatedAt"`
		Position             int         `json:"position"`
		StoryPoints          interface{} `json:"storyPoints"`
		RemainingTime        interface{} `json:"remainingTime"`
		CustomField1         int         `json:"customField1"`
		Embedded             Embedded    `json:"_embedded"`
		Links                struct {
			Attachments struct {
				Href string `json:"href"`
			} `json:"attachments"`
			AddAttachment struct {
				Href   string `json:"href"`
				Method string `json:"method"`
			} `json:"addAttachment"`
			Self struct {
				Href  string `json:"href"`
				Title string `json:"title"`
			} `json:"self"`
			Update struct {
				Href   string `json:"href"`
				Method string `json:"method"`
			} `json:"update"`
			Schema struct {
				Href string `json:"href"`
			} `json:"schema"`
			UpdateImmediately struct {
				Href   string `json:"href"`
				Method string `json:"method"`
			} `json:"updateImmediately"`
			Delete struct {
				Href   string `json:"href"`
				Method string `json:"method"`
			} `json:"delete"`
			LogTime struct {
				Href  string `json:"href"`
				Type  string `json:"type"`
				Title string `json:"title"`
			} `json:"logTime"`
			Move struct {
				Href  string `json:"href"`
				Type  string `json:"type"`
				Title string `json:"title"`
			} `json:"move"`
			Copy struct {
				Href  string `json:"href"`
				Title string `json:"title"`
			} `json:"copy"`
			Pdf struct {
				Href  string `json:"href"`
				Type  string `json:"type"`
				Title string `json:"title"`
			} `json:"pdf"`
			Atom struct {
				Href  string `json:"href"`
				Type  string `json:"type"`
				Title string `json:"title"`
			} `json:"atom"`
			AvailableRelationCandidates struct {
				Href  string `json:"href"`
				Title string `json:"title"`
			} `json:"availableRelationCandidates"`
			CustomFields struct {
				Href  string `json:"href"`
				Type  string `json:"type"`
				Title string `json:"title"`
			} `json:"customFields"`
			ConfigureForm struct {
				Href  string `json:"href"`
				Type  string `json:"type"`
				Title string `json:"title"`
			} `json:"configureForm"`
			Activities struct {
				Href string `json:"href"`
			} `json:"activities"`
			AvailableWatchers struct {
				Href string `json:"href"`
			} `json:"availableWatchers"`
			Relations struct {
				Href string `json:"href"`
			} `json:"relations"`
			Revisions struct {
				Href string `json:"href"`
			} `json:"revisions"`
			Watchers struct {
				Href string `json:"href"`
			} `json:"watchers"`
			AddWatcher struct {
				Href    string `json:"href"`
				Method  string `json:"method"`
				Payload struct {
					User struct {
						Href string `json:"href"`
					} `json:"user"`
				} `json:"payload"`
				Templated bool `json:"templated"`
			} `json:"addWatcher"`
			RemoveWatcher struct {
				Href      string `json:"href"`
				Method    string `json:"method"`
				Templated bool   `json:"templated"`
			} `json:"removeWatcher"`
			AddRelation struct {
				Href   string `json:"href"`
				Method string `json:"method"`
				Title  string `json:"title"`
			} `json:"addRelation"`
			AddChild struct {
				Href   string `json:"href"`
				Method string `json:"method"`
				Title  string `json:"title"`
			} `json:"addChild"`
			ChangeParent struct {
				Href   string `json:"href"`
				Method string `json:"method"`
				Title  string `json:"title"`
			} `json:"changeParent"`
			AddComment struct {
				Href   string `json:"href"`
				Method string `json:"method"`
				Title  string `json:"title"`
			} `json:"addComment"`
			PreviewMarkup struct {
				Href   string `json:"href"`
				Method string `json:"method"`
			} `json:"previewMarkup"`
			TimeEntries struct {
				Href  string `json:"href"`
				Type  string `json:"type"`
				Title string `json:"title"`
			} `json:"timeEntries"`
			Ancestors []interface{} `json:"ancestors"`
			Category  struct {
				Href interface{} `json:"href"`
			} `json:"category"`
			Type struct {
				Href  string `json:"href"`
				Title string `json:"title"`
			} `json:"type"`
			Priority struct {
				Href  string `json:"href"`
				Title string `json:"title"`
			} `json:"priority"`
			Project struct {
				Href  string `json:"href"`
				Title string `json:"title"`
			} `json:"project"`
			Status struct {
				Href  string `json:"href"`
				Title string `json:"title"`
			} `json:"status"`
			Author struct {
				Href  string `json:"href"`
				Title string `json:"title"`
			} `json:"author"`
			Responsible struct {
				Href interface{} `json:"href"`
			} `json:"responsible"`
			Assignee struct {
				Href  string `json:"href"`
				Title string `json:"title"`
			} `json:"assignee"`
			Version struct {
				Href  string `json:"href"`
				Title string `json:"title"`
			} `json:"version"`
			Parent struct {
				Uncacheable bool        `json:"uncacheable"`
				Href        interface{} `json:"href"`
				Title       interface{} `json:"title"`
			} `json:"parent"`
			CustomActions []interface{} `json:"customActions"`
		} `json:"_links"`
	} `json:"work_package"`
}

type TestJson struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type SlackPayload struct {
	AssigneeName  string `json:"name"`
	ID            int    `json:"issue_id"`
	Desc          string `json:"description"`
	IssuePriority string `json:"priority"`
}
