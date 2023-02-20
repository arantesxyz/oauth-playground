package opl_test

import (
	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

var observability_scenario_admins = []*rts.RelationTuple{
	//-------- create Org with Admin ---------
	// Organization: main
	// an organization is created with an initial admin
	{
		Namespace: "Organization",
		Object:    "main",
		Relation:  "admins",
		Subject: rts.NewSubjectSet(
			"Admin",
			"david",
			"",
		),
	},
	{
		Namespace: "Organization",
		Object:    "main",
		Relation:  "admins",
		Subject: rts.NewSubjectSet(
			"Admin",
			"hans",
			"",
		),
	},
}

var observability_scenario_users_groups = []*rts.RelationTuple{
	// -------- create empty Users ---------
	// empty user object is created so that we can check if a new user is allowed to be created by sombody
	// a relation to the main organization's admins is created
	{
		Namespace: "User",
		Object:    "",
		Relation:  "admins",
		Subject: rts.NewSubjectSet(
			"Organization",
			"main",
			"admins",
		),
	},
	// -------- create empty Groups ---------
	// empty group object is created so that we can check if a new group is allowed to be created by sombody
	// a relation to the main organization's admins is created
	{
		Namespace: "Group",
		Object:    "",
		Relation:  "admins",
		Subject: rts.NewSubjectSet(
			"Organization",
			"main",
			"admins",
		),
	},
	// -------- create a user in the AllUser group ---------
	// this group purposfully does not have an admin relation to the main organization
	// this is to test that admins can't accidentally delete the AllUsers group
	// when a new user is to be created this is done by adding the user to the AllUsers group
	// this is so that only subject sets are used to create new users
	{
		Namespace: "Group",
		Object:    "AllUsers",
		Relation:  "members",
		Subject: rts.NewSubjectSet(
			"User",
			"sam",
			"",
		),
	},
	{
		Namespace: "Group",
		Object:    "AllUsers",
		Relation:  "members",
		Subject: rts.NewSubjectSet(
			"User",
			"aaron",
			"",
		),
	},
	{
		Namespace: "Group",
		Object:    "AllUsers",
		Relation:  "members",
		Subject: rts.NewSubjectSet(
			"User",
			"nick",
			"",
		),
	},
	{
		Namespace: "Group",
		Object:    "AllUsers",
		Relation:  "members",
		Subject: rts.NewSubjectSet(
			"User",
			"cris",
			"",
		),
	},
	// -------- Add relations to the org admins for all users ---------
	{
		Namespace: "User",
		Object:    "sam",
		Relation:  "admins",
		Subject: rts.NewSubjectSet(
			"Organization",
			"main",
			"admins",
		),
	},
	{
		Namespace: "User",
		Object:    "aaron",
		Relation:  "admins",
		Subject: rts.NewSubjectSet(
			"Organization",
			"main",
			"admins",
		),
	},
	{
		Namespace: "User",
		Object:    "nick",
		Relation:  "admins",
		Subject: rts.NewSubjectSet(
			"Organization",
			"main",
			"admins",
		),
	},
	{
		Namespace: "User",
		Object:    "cris",
		Relation:  "admins",
		Subject: rts.NewSubjectSet(
			"Organization",
			"main",
			"admins",
		),
	},
	// -------- create actual group and Users ---------
	{
		Namespace: "Group",
		Object:    "MainCluster",
		Relation:  "admins",
		Subject: rts.NewSubjectSet(
			"Organization",
			"main",
			"admins",
		),
	},
	{
		Namespace: "Group",
		Object:    "FirstWorkloadCluster",
		Relation:  "admins",
		Subject: rts.NewSubjectSet(
			"Organization",
			"main",
			"admins",
		),
	},
	// -------- add users to group --------
	{
		Namespace: "Group",
		Object:    "MainCluster",
		Relation:  "members",
		Subject: rts.NewSubjectSet(
			"User",
			"sam",
			"",
		),
	},
	{
		Namespace: "Group",
		Object:    "FirstWorkloadCluster",
		Relation:  "members",
		Subject: rts.NewSubjectSet(
			"User",
			"aaron",
			"",
		),
	},
}

var observability_scenario_clients_tenants = []*rts.RelationTuple{
	// -------- create empty OAuth2Client ---------
	// empty OAuth2Client object is created so that we can check if a new OAuth2Client is allowed to be created by sombody
	// a relation to the main organization's admins is created
	{
		Namespace: "OAuth2Client",
		Object:    "",
		Relation:  "admins",
		Subject: rts.NewSubjectSet(
			"Organization",
			"main",
			"admins",
		),
	},
	// -------- create empty ObservabilityTenant ---------
	// empty ObservabilityTenant object is created so that we can check if a new ObservabilityTenant is allowed to be created by sombody
	// a relation to the main organization's admins is created
	{
		Namespace: "ObservabilityTenant",
		Object:    "",
		Relation:  "admins",
		Subject: rts.NewSubjectSet(
			"Organization",
			"main",
			"admins",
		),
	},
	// -------- create actual OAuth2Client and  add loginBindings for Users and Groups ---------
	{
		Namespace: "OAuth2Client",
		Object:    "MainClusterGrafana",
		Relation:  "admins",
		Subject: rts.NewSubjectSet(
			"Organization",
			"main",
			"admins",
		),
	},
	{
		Namespace: "OAuth2Client",
		Object:    "MainClusterGrafana",
		Relation:  "loginBindings",
		Subject: rts.NewSubjectSet(
			"Group",
			"MainCluster",
			"members",
		),
	},
	{
		Namespace: "OAuth2Client",
		Object:    "MainClusterGrafana",
		Relation:  "loginBindings",
		Subject: rts.NewSubjectSet(
			"User",
			"nick",
			"",
		),
	},
	// -------- create actual OAuth2Client for an agent sending data ---------
	{
		Namespace: "OAuth2Client",
		Object:    "MainClusterAgent",
		Relation:  "admins",
		Subject: rts.NewSubjectSet(
			"Organization",
			"main",
			"admins",
		),
	},
	// -------- create actual ObservabilityTenant and add loginBindings for Users and Groups ---------
	{
		Namespace: "ObservabilityTenant",
		Object:    "MainCluster",
		Relation:  "admins",
		Subject: rts.NewSubjectSet(
			"Organization",
			"main",
			"admins",
		),
	},
	{
		Namespace: "ObservabilityTenant",
		Object:    "MainCluster",
		Relation:  "viewers",
		Subject: rts.NewSubjectSet(
			"Group",
			"MainCluster",
			"members",
		),
	},
	// { // while technically possible and it would work, it isn't the intended workflow
	// 	Namespace: "ObservabilityTenant",
	// 	Object:    "MainCluster",
	// 	Relation:  "viewers",
	// 	Subject: rts.NewSubjectSet(
	// 		"OAuth2Client",
	// 		"MainClusterGrafana",
	// 		"",
	// 	),
	// },
	{
		Namespace: "ObservabilityTenant",
		Object:    "MainCluster",
		Relation:  "viewers",
		Subject: rts.NewSubjectSet(
			"OAuth2Client",
			"MainClusterAgent",
			"",
		),
	},
	{
		Namespace: "ObservabilityTenant",
		Object:    "MainCluster",
		Relation:  "editors",
		Subject: rts.NewSubjectSet(
			"User",
			"nick",
			"",
		),
	},
	{
		Namespace: "ObservabilityTenant",
		Object:    "FirstWorkloadCluster",
		Relation:  "admins",
		Subject: rts.NewSubjectSet(
			"Organization",
			"main",
			"admins",
		),
	},
	{
		Namespace: "ObservabilityTenant",
		Object:    "FirstWorkloadCluster",
		Relation:  "viewers",
		Subject: rts.NewSubjectSet(
			"Group",
			"FirstWorkloadCluster",
			"members",
		),
	},
	{
		Namespace: "ObservabilityTenant",
		Object:    "FirstWorkloadCluster",
		Relation:  "viewers",
		Subject: rts.NewSubjectSet(
			"Group",
			"MainCluster",
			"members",
		),
	},
	{
		Namespace: "ObservabilityTenant",
		Object:    "FirstWorkloadCluster",
		Relation:  "editors",
		Subject: rts.NewSubjectSet(
			"User",
			"cris",
			"",
		),
	},
}
