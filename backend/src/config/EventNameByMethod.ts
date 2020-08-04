export const eventNameByMethod: EventNameByMethod = {
  // ORGANIZATION
  CreateOrganization: 'organizationCreatedEvent',
  UnregisterOrganizationByID: 'organizationUnregisteredEvent',
  AddPublishedPatch: 'patchAddedToOrganization',
  UpdateOrganizationByID: 'organizationUpdatedEvent',

};

export type EventNameByMethod = {[methodName: string]: string};
