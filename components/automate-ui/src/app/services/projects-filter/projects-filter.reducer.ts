import { set, pipe, find } from 'lodash/fp';
import { EntityStatus } from 'app/entities/entities';
import {
  ProjectsFilterActions,
  ProjectsFilterActionTypes
} from './projects-filter.actions';

export interface ProjectsFilterOption {
  label: string;
  value: string;
  checked: boolean;
}

export interface ProjectsFilterState {
  options: ProjectsFilterOption[];
  optionsLoadingStatus: EntityStatus;
  selectionLabel: string;
  selectionCount: number;
  selectionCountVisible: boolean;
  selectionCountActive: boolean;
  dropdownCaretVisible: boolean;
}

export const projectsFilterInitialState: ProjectsFilterState = {
  options: [],
  optionsLoadingStatus: EntityStatus.notLoaded,
  selectionLabel: 'All resources',
  selectionCount: 0,
  selectionCountVisible: false,
  selectionCountActive: false,
  dropdownCaretVisible: false
};

export function projectsFilterReducer(
  state: ProjectsFilterState = projectsFilterInitialState,
  action: ProjectsFilterActions): ProjectsFilterState {

  switch (action.type) {

    case ProjectsFilterActionTypes.LOAD_OPTIONS: {
      return set('optionsLoadingStatus', EntityStatus.loading, state);
    }

    case ProjectsFilterActionTypes.LOAD_OPTIONS_SUCCESS: {
      const sortedOptions = action.payload.filter(o => o.value !== UNASSIGNED_PROJECT).sort(
          (a, b) => {
            const opts = { numeric: true, sensitivity: 'base' };
            return a.label.localeCompare(b.label, undefined, opts)
              || a.label.localeCompare(b.label, undefined, { numeric: true });
          }
        );
      const unassignedProject = find(['value', UNASSIGNED_PROJECT], action.payload);
      if (unassignedProject) {
        sortedOptions.push(unassignedProject);
      }
      return pipe(
        set('options', sortedOptions),
        set('optionsLoadingStatus', EntityStatus.loadingSuccess),
        set('selectionLabel', selectionLabel(action.payload)),
        set('selectionCount', selectionCount(action.payload)),
        set('selectionCountVisible', selectionCountVisible(action.payload)),
        set('selectionCountActive', selectionCountActive(action.payload)),
        set('dropdownCaretVisible', dropdownCaretVisible(action.payload))
      )(state) as ProjectsFilterState;
    }

    case ProjectsFilterActionTypes.LOAD_OPTIONS_FAILURE: {
      return set('optionsLoadingStatus', EntityStatus.loadingFailure, state);
    }

    case ProjectsFilterActionTypes.SAVE_OPTIONS: {
      return pipe(
        set('options', action.payload),
        set('selectionLabel', selectionLabel(action.payload)),
        set('selectionCount', selectionCount(action.payload)),
        set('selectionCountVisible', selectionCountVisible(action.payload)),
        set('selectionCountActive', selectionCountActive(action.payload)),
        set('dropdownCaretVisible', dropdownCaretVisible(action.payload))
      )(state) as ProjectsFilterState;
    }
  }

  return state;
}

const UNASSIGNED_PROJECT = '(unassigned)';

function selectionLabel(options: ProjectsFilterOption[]): string {
  const checkedOptions = options.filter(o => o.checked);
  const projectOptions = options.filter(o => o.value !== UNASSIGNED_PROJECT);
  const checkedProjects = projectOptions.filter(o => o.checked);
  const hasOneOption = options.length === 1;
  const hasOneChecked = checkedOptions.length === 1;
  const hasOneProjectChecked = checkedProjects.length === 1;
  const hasNoProjectsChecked = checkedProjects.length === 0;
  const hasAllProjectsChecked = checkedProjects.length === projectOptions.length;
  const hasSomeProjectsChecked = checkedProjects.length > 1 && !hasAllProjectsChecked;
  const hasUnassigned = options.filter(o => o.value === UNASSIGNED_PROJECT).length === 1;
  const hasUnassignedChecked =
    options.filter(o => o.value === UNASSIGNED_PROJECT && o.checked).length === 1;

  if (hasOneOption) {
    return options[0].label;
  }

  if (hasAllProjectsChecked && hasUnassignedChecked) {
    return 'All resources';
  }

  if (hasNoProjectsChecked && hasUnassigned && !hasUnassignedChecked) {
    return 'All resources';
  }

  if (hasOneChecked || hasOneProjectChecked) {
    return checkedOptions[0].label;
  }

  if (hasSomeProjectsChecked) {
    return 'Multiple projects';
  }

  return 'All projects';
}

function selectionCount(options: ProjectsFilterOption[]): number {
  const checkedProjects = options.filter(o => o.checked && o.value !== UNASSIGNED_PROJECT);
  return checkedProjects.length > 0 ? checkedProjects.length : options.length;
}

function selectionCountVisible(options: ProjectsFilterOption[]): boolean {
  const hasOnlyOneOption = options.length === 1;
  if (hasOnlyOneOption) {
    return false;
  }

  const checkedOptions = options.filter(o => o.checked);
  const projectOptions = options.filter(o => o.value !== UNASSIGNED_PROJECT);
  const checkedProjects = projectOptions.filter(o => o.checked);
  const hasOnlyProjects = projectOptions.length === options.length;
  const hasNoneChecked = checkedOptions.length === 0;
  const hasAllChecked = checkedOptions.length === options.length;
  const hasSomeProjectsChecked = checkedProjects.length > 1;

  return hasOnlyProjects ?
    hasSomeProjectsChecked || hasNoneChecked :
    hasSomeProjectsChecked && !hasAllChecked;
}

function selectionCountActive(options: ProjectsFilterOption[]): boolean {
  const checkedOptions = options.filter(o => o.checked);
  return checkedOptions.length > 0;
}

function dropdownCaretVisible(options: ProjectsFilterOption[]): boolean {
  const hasOnlyOneOption = options.length === 1;
  return !hasOnlyOneOption;
}
