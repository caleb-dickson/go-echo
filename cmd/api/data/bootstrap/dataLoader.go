package bootstrap

import (
	"fmt"
	"go-ng/cmd/api/data/model"
	"go-ng/cmd/api/data/model/entity"
)

func LoadPermissions() {

	var permissions []entity.Permission

	foundPermissions := model.DB.Find(&permissions)
	rows := foundPermissions.RowsAffected
	fmt.Printf("<----------|  %v Permissions in the database   |---------->\n", rows)

	permsPresent := foundPermissions.RowsAffected == 12

	if !permsPresent {
		var perms []entity.Permission

		viewUsers := entity.Permission{Name: "view_users"}
		editUsers := entity.Permission{Name: "edit_users"}
		deleteUsers := entity.Permission{Name: "delete_users"}

		viewRoles := entity.Permission{Name: "view_roles"}
		editRoles := entity.Permission{Name: "edit_roles"}
		deleteRoles := entity.Permission{Name: "delete_roles"}

		viewProducts := entity.Permission{Name: "view_products"}
		editProducts := entity.Permission{Name: "edit_products"}
		deleteProducts := entity.Permission{Name: "delete_products"}

		viewOrders := entity.Permission{Name: "view_orders"}
		editOrders := entity.Permission{Name: "edit_orders"}
		deleteOrders := entity.Permission{Name: "delete_orders"}

		perms = append(
			perms,
			viewUsers,
			editUsers,
			deleteUsers,
			viewRoles,
			editRoles,
			deleteRoles,
			viewProducts,
			editProducts,
			deleteProducts,
			viewOrders,
			editOrders,
			deleteOrders,
		)

		err := model.DB.Create(&perms).Error
		fmt.Println("<=======||    Permissions were added to the database!    ||=======>")
		if err != nil {
			_ = fmt.Sprintf("Error creating permissions: %v", err)
		}
	}

	fmt.Println("<----------|   Permissions already present!   |---------->")
}
