#include <libuast/uast.h>

static const char *
read_str(const void *data, const char *prop)
{
  PyObject *node = (PyObject *)data;
  PyObject *attribute = PyObject_GetAttrString(node, prop);
  const char *a = PyUnicode_AsUTF8(attribute);

  return a;
}

static int read_len(const void *data, const char *prop)
{
  PyObject *node = (PyObject *)data;
  PyObject *children_obj = PyObject_GetAttrString(node, prop);
  PyObject *seq = PySequence_Fast(children_obj, "expected a sequence");
  return PySequence_Size(children_obj);
}

static const char *
get_internal_type(const void *node)
{
  return read_str(node, "internal_type");
}

static const char *get_token(const void *node)
{
  return read_str(node, "token");
}

static int get_children_size(const void *node)
{
  return read_len(node, "children");
}

static void *get_child(const void *data, int index)
{
  PyObject *node = (PyObject *)data;
  PyObject *children_obj = PyObject_GetAttrString(node, "children");
  PyObject *seq = PySequence_Fast(children_obj, "expected a sequence");
  return PyList_GET_ITEM(seq, index);
}

static int get_roles_size(const void *node)
{
  return read_len(node, "roles");
}

static uint16_t get_role(const void *node, int index)
{
  return 2;
}

static node_api *api;

void create_go_node_api()
{
  api = new_node_api((node_iface){
      .internal_type = get_internal_type,
      .token = get_token,
      .children_size = get_children_size,
      .children = get_child,
      .roles_size = get_roles_size,
      .roles = get_role,
  });
}
