defineClass("JPTableViewController", {
  tableView_didSelectRowAtIndexPath: function(tableView, indexPath) {
    var row = indexPath.row()
    if (self.dataSource().count() > row) {
      var content = self.dataSource().objectAtIndex(row);
      var ctrl = JPViewController.alloc().initWithContent(content);
      self.navigationController().pushViewController(ctrl);
    }
  },
}, {})